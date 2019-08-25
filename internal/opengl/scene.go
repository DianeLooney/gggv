package opengl

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	var err error
	if *fullscreen {
		s.Window, err = glfw.CreateWindow(mode.Width, mode.Height, "gvd", monitor, nil)
	} else {
		s.Window, err = glfw.CreateWindow(mode.Width/2, mode.Height/2, "gvd", nil, nil)
	}

	if err != nil {
		panic(err)
	}
	s.Window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	s.Projection = mgl32.Ortho(-1, 1, -1, 1, 0.1, 10)
	s.Camera = mgl32.LookAtV(mgl32.Vec3{0, 0, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.Disable(gl.CULL_FACE)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)

	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
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
	FShaderLocation string
	VShaderLocation string
	GLProgram       uint32
}

type Uniform struct {
	Name  string
	Value interface{}
}

func (u Uniform) Bind(program uint32) {
	uLoc := gl.GetUniformLocation(program, gl.Str(u.Name+"\x00"))
	switch v := u.Value.(type) {
	case float32:
		gl.Uniform1f(uLoc, v)
	}
}

func (s *Scene) AddSourceFFVideo(name, path string) {
	dec, err := ffmpeg.NewFileSampler(path)
	if err != nil {
		fmt.Printf("Error adding new FFVideoSource: %v\n", err)
		return
	}
	var t uint32
	gl.GenTextures(1, &t)
	gl.ActiveTexture(t)
	gl.BindTexture(gl.TEXTURE_2D, t)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
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
	gl.GenFramebuffers(1, &sh.fbo)
	gl.BindFramebuffer(gl.FRAMEBUFFER, sh.fbo)
	gl.GenTextures(1, &sh.texture)
	gl.BindTexture(gl.TEXTURE_2D, sh.texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, s.Width, s.Height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, sh.texture, 0)

	gl.GenRenderbuffers(1, &sh.rbo)

	gl.BindRenderbuffer(gl.RENDERBUFFER, sh.rbo)

	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, s.Width, s.Height)
	gl.BindRenderbuffer(gl.RENDERBUFFER, 0)
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, sh.rbo)
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

func (s *Scene) LoadProgram(name, vertShaderLocation, fragShaderFlocation string) (err error) {
	data, err := ioutil.ReadFile(vertShaderLocation)
	if err != nil {
		return err
	}
	vertexShader, err := compileShader(string(data)+"\x00", gl.VERTEX_SHADER)
	if err != nil {
		return err
	}
	data, err = ioutil.ReadFile(fragShaderFlocation)
	if err != nil {
		return err
	}
	fragmentShader, err := compileShader(string(data)+"\x00", gl.FRAGMENT_SHADER)
	if err != nil {
		fmt.Println("Compile error:", err)
		return err
	}

	program := gl.CreateProgram()
	p := Program{
		FShaderLocation: fragShaderFlocation,
		VShaderLocation: vertShaderLocation,
		GLProgram:       program,
	}

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		//gl.DeleteProgram(program)
		return fmt.Errorf("failed to link program: %v", log)
	}

	if old, ok := s.programs[name]; ok {
		fmt.Printf("Deleting %v: old '%v', new '%v'\n", name, old.GLProgram, p.GLProgram)
		gl.DeleteProgram(old.GLProgram)
	}
	s.programs[name] = p

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	gl.UseProgram(program)

	return nil
}

func (s *Scene) BindBuffers() {
	// Configure the vertex data
	gl.GenVertexArrays(1, &s.vao)
	gl.BindVertexArray(s.vao)

	gl.GenBuffers(1, &s.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)

	/*
		{ // previousFrame pipeline setup
			gl.GenFramebuffers(1, &s.prevFrameFBObj)
			gl.BindFramebuffer(gl.FRAMEBUFFER, s.prevFrameFBObj)
			gl.GenTextures(1, &s.prevFrameFBTex)
			gl.BindTexture(gl.TEXTURE_2D, s.prevFrameFBTex)
			gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, s.Width, s.Height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
			gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, s.prevFrameFBTex, 0)

			gl.GenRenderbuffers(1, &s.prevFrameRBObj)

			gl.BindRenderbuffer(gl.RENDERBUFFER, s.prevFrameRBObj)

			gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, s.Width, s.Height)
			gl.BindRenderbuffer(gl.RENDERBUFFER, 0)
			gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, s.prevFrameRBObj)
		}
	*/
}

func (s *Scene) TextureInit(name string) {
	if _, ok := s.textures[name]; ok {
		return
	}

	var t uint32
	gl.GenTextures(1, &t)
	gl.ActiveTexture(t)
	gl.BindTexture(gl.TEXTURE_2D, t)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	s.textures[name] = t
}

func (s *Scene) RebindTexture(name string, width, height int, img []uint8) {
	t, ok := s.textures[name]
	if !ok {
		fmt.Printf("Unrecognized texture name %v\n", name)
		return
	}
	gl.ActiveTexture(t)
	gl.BindTexture(gl.TEXTURE_2D, t)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(width),
		int32(height),
		0,
		gl.RGB,
		gl.UNSIGNED_BYTE,
		gl.Ptr(&img[0]))
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func (s *Scene) Draw() {
	gl.UseProgram(s.programs["default"].GLProgram)

	s.time = float32(time.Since(tStart)) / NANOSTOSEC
	gl.BindVertexArray(s.vao)

	ord, err := Order("window", s.sources)
	if err != nil {
		fmt.Println("Error occurred while ordering sources for render:", err)
		return
	}
	for _, source := range ord {
		s.sources[source].Render(s)
	}
	{
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		program := s.programs["final"].GLProgram
		gl.UseProgram(program)

		src := s.sources["window"]
		//gl.BindTexture(gl.TEXTURE_2D, src.Texture())

		if shader, ok := src.(*ShaderSource); ok {
			for i, name := range shader.sources {
				if name == "" {
					continue
				}
				source := s.sources[name]
				gl.ActiveTexture(gl.TEXTURE0 + uint32(i))
				gl.BindTexture(gl.TEXTURE_2D, source.Texture())

				switch src := source.(type) {
				case *FFVideoSource:
					x := gl.GetUniformLocation(program, gl.Str(fmt.Sprintf("tex%vwidth\x00", i)))
					gl.Uniform1f(x, float32(src.width))

					x = gl.GetUniformLocation(program, gl.Str(fmt.Sprintf("tex%vheight\x00", i)))
					gl.Uniform1f(x, float32(src.width))
				}
			}
		}
		s.bindCommonUniforms(program)

		projectionMat := mgl32.Ortho(-1, 1, -1, 1, 0.1, 10)
		projection := gl.GetUniformLocation(program, gl.Str("projection\x00"))
		gl.UniformMatrix4fv(projection, 1, false, &projectionMat[0])

		//bind framebuffer texture
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BufferData(gl.ARRAY_BUFFER, len(staticVerts)*4, gl.Ptr(&staticVerts[0]), gl.STATIC_DRAW)
		gl.DrawArrays(gl.TRIANGLES, 0, 2*3)

		//draw output wuad
	}

	s.Window.SwapBuffers()
	fps.Next()
	glfw.PollEvents()
}

func (s *Scene) bindCommonUniforms(program uint32) {
	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	projection := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projection, 1, false, &s.Projection[0])

	camera := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(camera, 1, false, &s.Camera[0])

	for i := int32(0); i < SHADER_TEXTURE_COUNT; i++ {
		tex := gl.GetUniformLocation(program, gl.Str(fmt.Sprintf("tex%v\x00", i)))
		gl.Uniform1i(tex, i)
	}

	timeU := gl.GetUniformLocation(program, gl.Str("time\x00"))
	gl.Uniform1f(timeU, s.time)

	fpsU := gl.GetUniformLocation(program, gl.Str("fps\x00"))
	gl.Uniform1f(fpsU, float32(fps.LastSec()))

	renderTime := gl.GetUniformLocation(program, gl.Str("renderTime\x00"))
	gl.Uniform1f(renderTime, float32(fps.FrameDuration())/NANOSTOSEC)

	x, y := s.Window.GetCursorPos()
	cursorXU := gl.GetUniformLocation(program, gl.Str("cursorX\x00"))
	gl.Uniform1f(cursorXU, float32(x))

	cursorYU := gl.GetUniformLocation(program, gl.Str("cursorY\x00"))
	gl.Uniform1f(cursorYU, float32(y))

	windowWidth, windowHeight := s.Window.GetSize()
	windowWidthU := gl.GetUniformLocation(program, gl.Str("windowWidth\x00"))
	gl.Uniform1f(windowWidthU, float32(windowWidth))

	windowHeightU := gl.GetUniformLocation(program, gl.Str("windowHeight\x00"))
	gl.Uniform1f(windowHeightU, float32(windowHeight))
}

func (s *Scene) ReloadPrograms() {
	work := make([]Program, len(s.programs))
	names := make([]string, len(s.programs))
	i := 0
	for name, program := range s.programs {
		work[i] = program
		names[i] = name
		i++
	}
	for i, program := range work {
		s.LoadProgram(names[i], program.VShaderLocation, program.FShaderLocation)
	}
}
