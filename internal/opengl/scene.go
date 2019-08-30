package opengl

import (
	"flag"
	"fmt"
	"io/ioutil"
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

var tStart = time.Now()

var borderless = flag.Bool("borderless", true, "Hide borders")
var fullscreen = flag.Bool("fullscreen", true, "Start in fullscreen mode")

var vsync = flag.Bool("vsync", true, "Enable/Disable vsync")

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
	}

	if *borderless {
		glfw.WindowHint(glfw.Decorated, glfw.False)
	}

	var err error
	if *fullscreen {
		s.Window, err = glfw.CreateWindow(mode.Width, mode.Height, "gvd", nil, nil)
	} else {
		s.Window, err = glfw.CreateWindow(mode.Width/2, mode.Height/2, "gvd", nil, nil)
	}

	if err != nil {
		panic(err)
	}
	s.Window.MakeContextCurrent()

	// Initialize Glow
	if err := carbon.Init(); err != nil {
		panic(err)
	}

	s.Projection = mgl32.Ortho(-1, 1, -1, 1, 0.1, 10)
	s.Camera = mgl32.LookAtV(mgl32.Vec3{0, 0, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})

	// Configure global settings
	carbon.Enable(carbon.DEPTH_TEST)
	carbon.Disable(carbon.CULL_FACE)
	carbon.Enable(carbon.BLEND)
	carbon.BlendFunc(carbon.SRC_ALPHA, carbon.ONE_MINUS_SRC_ALPHA)
	carbon.PixelStorei(carbon.UNPACK_ALIGNMENT, 1)

	carbon.DepthFunc(carbon.LESS)
	carbon.ClearColor(0, 0, 0, 1)
	carbon.Clear(carbon.COLOR_BUFFER_BIT | carbon.DEPTH_BUFFER_BIT)
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

	Camera     mgl32.Mat4
	Projection mgl32.Mat4

	Width  int32
	Height int32

	time float32

	programs map[string]Program
	textures map[string]uint32

	sources map[SourceName]Source
}

type SourceKind string

type SourceName string

type Source interface {
	Kind() SourceKind
	Name() SourceName
	Children() []SourceName
	Render(scene *Scene)
	Texture() uint32
}

type Program struct {
	GLProgram uint32
}

type Uniform struct {
	Name  string
	Value interface{}
}

func (s *Scene) AddSourceFFVideo(name, path string) {
	dec, err := ffmpeg.NewFileSampler(path)
	if err != nil {
		logs.Error("Error adding new FFVideoSource", err)
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
	s.sources[SourceName(name)] = &FFVideoSource{
		name:    SourceName(name),
		decoder: dec,
		texture: t,
	}
}

func (s *Scene) AddSourceShader(name string) {
	s.LoadProgram(name, "shaders/vert/default.glsl", "shaders/frag/default.glsl")
	sh := ShaderSource{
		name:       SourceName(name),
		uniforms:   make(map[string]Uniform),
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
	vShader, err := ioutil.ReadFile("shaders/vert/window.glsl")
	if err != nil {
		panic(err)
	}
	fShader, err := ioutil.ReadFile("shaders/frag/window.glsl")
	if err != nil {
		panic(err)
	}

	if err := s.LoadProgram("window", string(vShader), string(fShader)); err != nil {
		panic(err)
	}
	s.sources[SourceName("window")] = &ShaderSource{
		name:       SourceName("window"),
		uniforms:   make(map[string]Uniform),
		p:          "window",
		flipOutput: false,
	}
}

func (s *Scene) SetUniform(layer, name string, value interface{}) {
	src, ok := s.sources[SourceName(layer)]
	if !ok {
		return
	}
	if shader, ok := src.(*ShaderSource); ok {
		shader.uniforms[name] = Uniform{name, value}
	}
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

func (s *Scene) LoadProgram(name, vShader, fShader string) (err error) {
	vertexShader, err := compileShader(vShader+"\x00", carbon.VERTEX_SHADER)
	if err != nil {
		return err
	}
	fragmentShader, err := compileShader(fShader+"\x00", carbon.FRAGMENT_SHADER)
	if err != nil {
		return err
	}

	program := carbon.CreateProgram()
	p := Program{
		GLProgram: program,
	}

	carbon.AttachShader(program, vertexShader)
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

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := carbon.CreateShader(shaderType)

	csources, free := carbon.Strs(source)
	carbon.ShaderSource(shader, 1, csources, nil)
	free()
	carbon.CompileShader(shader)

	var status int32
	carbon.GetShaderiv(shader, carbon.COMPILE_STATUS, &status)
	if status == carbon.FALSE {
		var logLength int32
		carbon.GetShaderiv(shader, carbon.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		carbon.GetShaderInfoLog(shader, logLength, nil, carbon.Str(log))

		return 0, errors.GLShaderCompile(log, source)
	}

	return shader, nil
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
		logs.Error(errors.SceneRenderOrder(err))
		return
	}

	for _, source := range ord {
		s.sources[source].Render(s)
	}
	windowSrc.Render(s)

	s.Window.SwapBuffers()
	fps.Next()
	glfw.PollEvents()
}

func (s *Scene) bindCommonUniforms(program uint32) {
	vertAttrib := uint32(carbon.GetAttribLocation(program, carbon.Str("vert\x00")))
	carbon.EnableVertexAttribArray(vertAttrib)
	carbon.VertexAttribPointer(vertAttrib, 3, carbon.FLOAT, false, 5*4, carbon.PtrOffset(0))

	texCoordAttrib := uint32(carbon.GetAttribLocation(program, carbon.Str("vertTexCoord\x00")))
	carbon.EnableVertexAttribArray(texCoordAttrib)
	carbon.VertexAttribPointer(texCoordAttrib, 2, carbon.FLOAT, false, 5*4, carbon.PtrOffset(3*4))

	carbon.Uniform(program, "projection", s.Projection)
	carbon.Uniform(program, "camera", s.Camera)

	for i := 0; i < SHADER_TEXTURE_COUNT; i++ {
		carbon.UniformTex(program, fmt.Sprintf("tex%v", i), int32(i))
	}

	carbon.Uniform(program, "time", s.time)

	fpsU := carbon.GetUniformLocation(program, carbon.Str("fps\x00"))
	carbon.Uniform1f(fpsU, float32(fps.LastSec()))

	renderTime := carbon.GetUniformLocation(program, carbon.Str("renderTime\x00"))
	carbon.Uniform1f(renderTime, float32(fps.FrameDuration())/NANOSTOSEC)

	x, y := s.Window.GetCursorPos()
	carbon.Uniform(program, "cursorX", x)
	carbon.Uniform(program, "cursorY", y)

	windowWidth, windowHeight := s.Window.GetSize()
	carbon.Uniform(program, "windowWidth", windowWidth)
	carbon.Uniform(program, "windowHeight", windowHeight)
}
