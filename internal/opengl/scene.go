package opengl

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/dianelooney/gggv/internal/carbon"
	"github.com/dianelooney/gggv/internal/errors"
	"github.com/dianelooney/gggv/internal/ffmpeg"
	"github.com/dianelooney/gggv/internal/logs"

	"github.com/dianelooney/gggv/internal/fps"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const NANOSTOSEC = 1000000000

var tStart = time.Now().Add(-1 * time.Second) // subtracted a second to enforce non-zero times

var borderless = flag.Bool("borderless", false, "Hide borders")
var fullscreen = flag.Bool("fullscreen", false, "Start in fullscreen mode")

var vsync = flag.Bool("vsync", true, "Enable/Disable vsync")
var width = flag.Int("width", -1, "Set the window width")
var height = flag.Int("height", -1, "Set the window height")

func NewScene() *Scene {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	s := &Scene{
		programs: make(map[string]Program),
		textures: make(map[string]uint32),
		sources:  make(map[SourceName]Source),
		uniforms: make(map[string]BindUniformer),
	}

	if *borderless {
		glfw.WindowHint(glfw.Decorated, glfw.False)
	}

	var err error
	if *fullscreen {
		s.Window, err = glfw.CreateWindow(mode.Width, mode.Height, "gggv", nil, nil)
	} else {
		var w = *width
		var h = *height
		if h < 0 {
			h = mode.Height / 2
		}
		if w < 0 {
			w = mode.Width / 2
		}
		s.Window, err = glfw.CreateWindow(w, h, "gggv", nil, nil)
	}

	if err != nil {
		log.Fatalf("Unable to create window: %v", err)
	}

	s.Window.MakeContextCurrent()

	// Initialize Glow
	if err := carbon.Init(); err != nil {
		log.Fatalf("Unable to initialize glow: %v", err)
	}

	s.Camera = mgl32.LookAtV(mgl32.Vec3{0, 0, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})

	// Configure global settings
	carbon.PixelStorei(carbon.UNPACK_ALIGNMENT, 1)

	carbon.ClearColor(0, 0, 0, 1)
	if *vsync {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}

	w, h := s.Window.GetFramebufferSize()
	s.Width, s.Height = int32(w), int32(h)

	return s
}

func finalizeScene(s *Scene) {
	defer glfw.Terminate()
}

type Scene struct {
	Window *glfw.Window

	vao uint32
	vbo uint32

	Camera mgl32.Mat4

	Width  int32
	Height int32

	time float32

	programs map[string]Program
	textures map[string]uint32
	uniforms map[string]BindUniformer

	sources map[SourceName]Source
}

type SourceName string

type Source interface {
	Name() SourceName
	Children() []SourceName
	Render(scene *Scene)
	SkipRender(scene *Scene)
	Texture() uint32
}

type Program struct {
	GLProgram uint32
}

type ValueUniform struct {
	Name  string
	Value interface{}
}

func (u ValueUniform) BindUniform(program uint32) {
	carbon.Uniform(program, u.Name, u.Value)
}

type ClockUniform struct {
	Name   string
	Offset time.Time
}

func (u ClockUniform) BindUniform(program uint32) {
	carbon.Uniform(program, u.Name, float32(time.Since(u.Offset))/NANOSTOSEC)
}

type BindUniformer interface {
	BindUniform(program uint32)
}

func (s *Scene) AddSourceFFT(name string) {
	s.sources[SourceName(name)] = NewPortaudio()
}

func (s *Scene) AddSourceFFVideo(name, path string) {
	reader := ffmpeg.NewTimer(ffmpeg.Buffer(ffmpeg.Loop(func() (ffmpeg.Reader, error) {
		return ffmpeg.NewReader(path)
	})))

	var t uint32
	carbon.GenTextures(1, &t)
	carbon.ActiveTexture(t)
	carbon.BindTexture(carbon.TEXTURE_2D, t)
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MIN_FILTER, carbon.LINEAR)
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MAG_FILTER, carbon.LINEAR)
	s.sources[SourceName(name)] = &FFVideoSource{
		name:    SourceName(name),
		decoder: reader,
		texture: t,
	}
}

func (s *Scene) AddSourceShader(name string) {
	sh := ShaderSource{
		name:       SourceName(name),
		uniforms:   make(map[string]BindUniformer),
		p:          name,
		flipOutput: true,
	}
	carbon.GenFramebuffers(1, &sh.fbo)
	carbon.BindFramebuffer(carbon.FRAMEBUFFER, sh.fbo)
	carbon.GenTextures(1, &sh.texture)
	carbon.BindTexture(carbon.TEXTURE_2D, sh.texture)
	carbon.TexImage2D(carbon.TEXTURE_2D, 0, carbon.RGB, s.Width, s.Height, 0, carbon.RGB, carbon.UNSIGNED_BYTE, nil)
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MIN_FILTER, carbon.LINEAR)
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MAG_FILTER, carbon.LINEAR)
	carbon.FramebufferTexture2D(carbon.FRAMEBUFFER, carbon.COLOR_ATTACHMENT0, carbon.TEXTURE_2D, sh.texture, 0)

	carbon.GenRenderbuffers(1, &sh.rbo)

	carbon.BindRenderbuffer(carbon.RENDERBUFFER, sh.rbo)

	carbon.RenderbufferStorage(carbon.RENDERBUFFER, carbon.DEPTH24_STENCIL8, s.Width, s.Height)
	carbon.BindRenderbuffer(carbon.RENDERBUFFER, 0)
	carbon.FramebufferRenderbuffer(carbon.FRAMEBUFFER, carbon.DEPTH_STENCIL_ATTACHMENT, carbon.RENDERBUFFER, sh.rbo)
	s.sources[SourceName(name)] = &sh
}

func (s *Scene) AddWindow() {
	s.sources[SourceName("window")] = &ShaderSource{
		name:     SourceName("window"),
		uniforms: make(map[string]BindUniformer),
		p:        "window",
	}
}

func (s *Scene) SetFFVideoTimescale(name string, timescale float64) {
	if src, ok := s.sources[SourceName(name)]; ok {
		if ffv, ok := src.(*FFVideoSource); ok {
			ffv.decoder.Timescale(timescale)
		}
	}
}

func (s *Scene) SetFFTScale(name string, scale float32) {
	if src, ok := s.sources[SourceName(name)]; ok {
		if ffv, ok := src.(*Portaudio); ok {
			ffv.scale = scale
		}
	}
}

func (s *Scene) SetShaderProgram(name, program string) {
	if src, ok := s.sources[SourceName(name)]; ok {
		if sh, ok := src.(*ShaderSource); ok {
			sh.p = program
		}
	}
}

func (s *Scene) SetSourceMinFilter(name, value string) {
	opt, ok := map[string]int32{
		"NEAREST":                carbon.NEAREST,
		"LINEAR":                 carbon.LINEAR,
		"NEAREST_MIPMAP_NEAREST": carbon.NEAREST_MIPMAP_NEAREST,
		"LINEAR_MIPMAP_NEAREST":  carbon.LINEAR_MIPMAP_NEAREST,
		"NEAREST_MIPMAP_LINEAR":  carbon.NEAREST_MIPMAP_LINEAR,
		"LINEAR_MIPMAP_LINEAR":   carbon.LINEAR_MIPMAP_LINEAR,
	}[value]
	if !ok {
		return
	}
	src := s.sources[SourceName(name)]
	carbon.ActiveTexture(src.Texture())
	//TODO: better error handling on all of these.
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MIN_FILTER, opt)
}
func (s *Scene) SetSourceMagFilter(name, value string) {
	opt, ok := map[string]int32{
		"NEAREST": carbon.NEAREST,
		"LINEAR":  carbon.LINEAR,
	}[value]
	if !ok {
		return
	}
	src := s.sources[SourceName(name)]
	carbon.ActiveTexture(src.Texture())
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MAG_FILTER, opt)
}
func (s *Scene) SetSourceWrapS(name, value string) {
	opt, ok := map[string]int32{
		"CLAMP_TO_EDGE":        carbon.CLAMP_TO_EDGE,
		"CLAMP_TO_BORDER":      carbon.CLAMP_TO_BORDER,
		"MIRRORED_REPEAT":      carbon.MIRRORED_REPEAT,
		"REPEAT":               carbon.REPEAT,
		"MIRROR_CLAMP_TO_EDGE": carbon.MIRROR_CLAMP_TO_EDGE,
	}[value]
	if !ok {
		return
	}
	src := s.sources[SourceName(name)]
	carbon.ActiveTexture(src.Texture())
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_WRAP_S, opt)
}
func (s *Scene) SetSourceWrapT(name, value string) {
	opt, ok := map[string]int32{
		"CLAMP_TO_EDGE":        carbon.CLAMP_TO_EDGE,
		"CLAMP_TO_BORDER":      carbon.CLAMP_TO_BORDER,
		"MIRRORED_REPEAT":      carbon.MIRRORED_REPEAT,
		"REPEAT":               carbon.REPEAT,
		"MIRROR_CLAMP_TO_EDGE": carbon.MIRROR_CLAMP_TO_EDGE,
	}[value]
	if !ok {
		return
	}
	src := s.sources[SourceName(name)]
	carbon.ActiveTexture(src.Texture())
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_WRAP_T, opt)
}

func (s *Scene) SetUniform(layer, name string, value interface{}) {
	src, ok := s.sources[SourceName(layer)]
	if !ok {
		return
	}
	if shader, ok := src.(*ShaderSource); ok {
		shader.uniforms[name] = ValueUniform{name, value}
	}
}

func (s *Scene) SetGlobalUniform(name string, value interface{}) {
	s.uniforms[name] = ValueUniform{name, value}
}

func (s *Scene) SetUniformClock(layer, name string, offset time.Time) {
	src, ok := s.sources[SourceName(layer)]
	if !ok {
		return
	}
	if shader, ok := src.(*ShaderSource); ok {
		shader.uniforms[name] = ClockUniform{name, offset}
	}
}

func (s *Scene) SetGlobalUniformClock(name string, offset time.Time) {
	s.uniforms[name] = ClockUniform{name, offset}
}

func (s *Scene) SetShaderInput(layer string, index int32, target string) {
	l, ok := s.sources[SourceName(layer)]
	if !ok {
		logs.Error("Attempted to set input on layer, but the layer was not found", layer)
		return
	}
	src, ok := l.(*ShaderSource)
	if !ok {
		logs.Error("Attempted to set input on layer, but the layer was not a shader", layer)
		return
	}
	src.sources[index] = SourceName(target)
	s.sources[SourceName(layer)] = src
}

func (s *Scene) LoadProgram(name, vShader, gShader, fShader string) (err error) {
	vertexShader, err := carbon.WrappedCompileShader(vShader+"\x00", carbon.VERTEX_SHADER)
	if err != nil {
		return err
	}
	geometryShader, err := carbon.WrappedCompileShader(gShader+"\x00", carbon.GEOMETRY_SHADER)
	if err != nil {
		return err
	}
	fragmentShader, err := carbon.WrappedCompileShader(fShader+"\x00", carbon.FRAGMENT_SHADER)
	if err != nil {
		return err
	}

	program := carbon.CreateProgram()
	p := Program{
		GLProgram: program,
	}

	carbon.AttachShader(program, vertexShader)
	carbon.AttachShader(program, geometryShader)
	carbon.AttachShader(program, fragmentShader)
	carbon.LinkProgram(program)

	var status int32
	carbon.GetProgramiv(program, carbon.LINK_STATUS, &status)
	if status == carbon.FALSE {
		var logLength int32
		carbon.GetProgramiv(program, carbon.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		carbon.GetProgramInfoLog(program, logLength, nil, carbon.Str(log))

		return errors.GLLinkProgram(name, log)
	}

	if old, ok := s.programs[name]; ok {
		carbon.DeleteProgram(old.GLProgram)
	}
	s.programs[name] = p

	carbon.BindFragDataLocation(program, 0, carbon.Str("outputColor\x00"))
	carbon.DeleteShader(vertexShader)
	carbon.DeleteShader(fragmentShader)

	return nil
}

func (s *Scene) BindBuffers() {
	carbon.GenVertexArrays(1, &s.vao)
	carbon.BindVertexArray(s.vao)

	carbon.GenBuffers(1, &s.vbo)
	carbon.BindBuffer(carbon.ARRAY_BUFFER, s.vbo)
}

func (s *Scene) TextureInit(name string) {
	if _, ok := s.textures[name]; ok {
		return
	}

	var t uint32
	carbon.GenTextures(1, &t)
	carbon.ActiveTexture(t)
	carbon.BindTexture(carbon.TEXTURE_2D, t)
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MIN_FILTER, carbon.LINEAR)
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MAG_FILTER, carbon.LINEAR)
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_WRAP_S, carbon.CLAMP_TO_EDGE)
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_WRAP_T, carbon.CLAMP_TO_EDGE)
	s.textures[name] = t
}

func (s *Scene) RebindTexture(name string, width, height int, img []uint8) {
	t, ok := s.textures[name]
	if !ok {
		logs.Error("Unrecognized texture name", name) // TODO: named error
		return
	}
	carbon.ActiveTexture(t)
	carbon.BindTexture(carbon.TEXTURE_2D, t)
	carbon.TexImage2D(
		carbon.TEXTURE_2D,
		0,
		carbon.RGBA,
		int32(width),
		int32(height),
		0,
		carbon.RGB,
		carbon.UNSIGNED_BYTE,
		carbon.Ptr(&img[0]))
}

func (s *Scene) Draw() {
	windowSrc, ok := s.sources["window"]
	if !ok {
		logs.Error(errors.SceneMissingWindowSource())
		return
	}

	s.time = float32(time.Since(tStart)) / NANOSTOSEC
	carbon.BindVertexArray(s.vao)

	ord, err := Order("window", s.sources)
	if err != nil {
		// logs.Error(errors.SceneRenderOrder(err))
		return
	}

	rendered := make(map[SourceName]bool, len(ord))
	for _, source := range ord {
		s.sources[source].Render(s)
		rendered[source] = true
	}
	for name, source := range s.sources {
		if !rendered[name] {
			source.SkipRender(s)
		}
	}
	windowSrc.Render(s)

	s.Window.SwapBuffers()
	fps.Next()
	glfw.PollEvents()
}
