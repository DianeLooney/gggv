package opengl

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dianelooney/gggv/internal/ffmpeg"

	"github.com/dianelooney/gggv/internal/fps"

	"github.com/go-gl/gl/all-core/gl"
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
type Program struct {
	GLProgram uint32
}

type Uniform struct {
	Name  string
	Value interface{}
}

func (u Uniform) Bind(program uint32) {
	uLoc := carbon.GetUniformLocation(program, carbon.Str(u.Name+"\x00"))
	switch v := u.Value.(type) {
	case float32:
		carbon.Uniform1f(uLoc, v)
	}
}

func (s *Scene) AddSourceFFVideo(name, path string) {
	dec, err := ffmpeg.NewFileSampler(path)
	if err != nil {
		fmt.Printf("Error adding new FFVideoSource: %v\n", err)
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
		name:     SourceName(name),
		uniforms: make(map[string]Uniform),
		p:        name,
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
		fmt.Printf("Attempted to set input on %v, but %v was not found.\n", layer, layer)
		fmt.Println(s.sources)
		return
	}
	fmt.Println(s.sources[SourceName(layer)])
	src, ok := l.(*ShaderSource)
	if !ok {
		fmt.Printf("Attempted to set input on %v, but %v was not a shader.\n", layer, layer)
		return
	}
	src.sources[index] = SourceName(target)
	s.sources[SourceName(layer)] = src
	fmt.Println(s.sources[SourceName(layer)])
}

func (s *Scene) LoadProgram(name, vShader, fShader string) (err error) {
	vertexShader, err := compileShader(vShader+"\x00", carbon.VERTEX_SHADER)
	if err != nil {
		return err
	}
	fragmentShader, err := compileShader(fShader+"\x00", carbon.FRAGMENT_SHADER)
	if err != nil {
		fmt.Println("Compile error:", err)
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

		//carbon.DeleteProgram(program)
		return fmt.Errorf("failed to link program: %v", log)
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
	// Configure the vertex data
	carbon.GenVertexArrays(1, &s.vao)
	carbon.BindVertexArray(s.vao)

	carbon.GenBuffers(1, &s.vbo)
	carbon.BindBuffer(carbon.ARRAY_BUFFER, s.vbo)

	/*
		{ // previousFrame pipeline setup
			carbon.GenFramebuffers(1, &s.prevFrameFBObj)
			carbon.BindFramebuffer(carbon.FRAMEBUFFER, s.prevFrameFBObj)
			carbon.GenTextures(1, &s.prevFrameFBTex)
			carbon.BindTexture(carbon.TEXTURE_2D, s.prevFrameFBTex)
			carbon.TexImage2D(carbon.TEXTURE_2D, 0, carbon.RGB, s.Width, s.Height, 0, carbon.RGB, carbon.UNSIGNED_BYTE, nil)
			carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MIN_FILTER, carbon.LINEAR)
			carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MAG_FILTER, carbon.LINEAR)
			carbon.FramebufferTexture2D(carbon.FRAMEBUFFER, carbon.COLOR_ATTACHMENT0, carbon.TEXTURE_2D, s.prevFrameFBTex, 0)

			carbon.GenRenderbuffers(1, &s.prevFrameRBObj)

			carbon.BindRenderbuffer(carbon.RENDERBUFFER, s.prevFrameRBObj)

			carbon.RenderbufferStorage(carbon.RENDERBUFFER, carbon.DEPTH24_STENCIL8, s.Width, s.Height)
			carbon.BindRenderbuffer(carbon.RENDERBUFFER, 0)
			carbon.FramebufferRenderbuffer(carbon.FRAMEBUFFER, carbon.DEPTH_STENCIL_ATTACHMENT, carbon.RENDERBUFFER, s.prevFrameRBObj)
		}
	*/
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
		fmt.Printf("Unrecognized texture name %v\n", name)
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

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func (s *Scene) Draw() {
	carbon.UseProgram(s.programs["default"].GLProgram)

	s.time = float32(time.Since(tStart)) / NANOSTOSEC
	carbon.BindVertexArray(s.vao)

	ord, err := Order("window", s.sources)
	if err != nil {
		fmt.Println("Error occurred while ordering sources for render:", err)
		return
	}
	for _, source := range ord {
		s.sources[source].Render(s)
	}
	{
		carbon.Clear(carbon.COLOR_BUFFER_BIT | carbon.DEPTH_BUFFER_BIT)

		program := s.programs["window"].GLProgram
		carbon.UseProgram(program)

		src := s.sources["window"]
		//carbon.BindTexture(carbon.TEXTURE_2D, src.Texture())

		if shader, ok := src.(*ShaderSource); ok {
			for i, name := range shader.sources {
				if name == "" {
					continue
				}
				source := s.sources[name]
				carbon.ActiveTexture(carbon.TEXTURE0 + uint32(i))
				carbon.BindTexture(carbon.TEXTURE_2D, source.Texture())

				switch src := source.(type) {
				case *FFVideoSource:
					x := carbon.GetUniformLocation(program, carbon.Str(fmt.Sprintf("tex%vwidth\x00", i)))
					carbon.Uniform1f(x, float32(src.width))

					x = carbon.GetUniformLocation(program, carbon.Str(fmt.Sprintf("tex%vheight\x00", i)))
					carbon.Uniform1f(x, float32(src.width))
				}
			}
		}
		s.bindCommonUniforms(program)

		projectionMat := mgl32.Ortho(-1, 1, -1, 1, 0.1, 10)
		projection := carbon.GetUniformLocation(program, carbon.Str("projection\x00"))
		carbon.UniformMatrix4fv(projection, 1, false, &projectionMat[0])

		//bind framebuffer texture
		carbon.ActiveTexture(carbon.TEXTURE0)
		carbon.BufferData(carbon.ARRAY_BUFFER, len(staticVerts)*4, carbon.Ptr(&staticVerts[0]), carbon.STATIC_DRAW)
		carbon.DrawArrays(carbon.TRIANGLES, 0, 2*3)

		//draw output wuad
	}

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

	projection := carbon.GetUniformLocation(program, carbon.Str("projection\x00"))
	carbon.UniformMatrix4fv(projection, 1, false, &s.Projection[0])

	camera := carbon.GetUniformLocation(program, carbon.Str("camera\x00"))
	carbon.UniformMatrix4fv(camera, 1, false, &s.Camera[0])

	for i := int32(0); i < SHADER_TEXTURE_COUNT; i++ {
		tex := carbon.GetUniformLocation(program, carbon.Str(fmt.Sprintf("tex%v\x00", i)))
		carbon.Uniform1i(tex, i)
	}

	timeU := carbon.GetUniformLocation(program, carbon.Str("time\x00"))
	carbon.Uniform1f(timeU, s.time)

	fpsU := carbon.GetUniformLocation(program, carbon.Str("fps\x00"))
	carbon.Uniform1f(fpsU, float32(fps.LastSec()))

	renderTime := carbon.GetUniformLocation(program, carbon.Str("renderTime\x00"))
	carbon.Uniform1f(renderTime, float32(fps.FrameDuration())/NANOSTOSEC)

	x, y := s.Window.GetCursorPos()
	cursorXU := carbon.GetUniformLocation(program, carbon.Str("cursorX\x00"))
	carbon.Uniform1f(cursorXU, float32(x))

	cursorYU := carbon.GetUniformLocation(program, carbon.Str("cursorY\x00"))
	carbon.Uniform1f(cursorYU, float32(y))

	windowWidth, windowHeight := s.Window.GetSize()
	windowWidthU := carbon.GetUniformLocation(program, carbon.Str("windowWidth\x00"))
	carbon.Uniform1f(windowWidthU, float32(windowWidth))

	windowHeightU := carbon.GetUniformLocation(program, carbon.Str("windowHeight\x00"))
	carbon.Uniform1f(windowHeightU, float32(windowHeight))
}
