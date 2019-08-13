package opengl

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"

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
		layers:   make(map[string]Layer),
		programs: make(map[string]Program),
		textures: make(map[string]uint32),
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

	Camera       mgl32.Mat4
	Projection   mgl32.Mat4
	ModelUniform int32
	TimeUniform  int32
	DepthUniform int32

	Width  int32
	Height int32

	prevPassFBObj uint32
	prevPassFBTex uint32
	prevPassRBObj uint32

	prevFrameFBObj uint32
	prevFrameFBTex uint32
	prevFrameRBObj uint32

	layers   map[string]Layer
	programs map[string]Program
	textures map[string]uint32
}
type Program struct {
	FShaderLocation string
	VShaderLocation string
	GLProgram       uint32
}
type Layer struct {
	Depth   float32
	Texture string
	Program string
}

type layers []Layer

func (l layers) Less(i, j int) bool {
	return l[i].Depth < l[j].Depth
}
func (l layers) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l layers) Len() int {
	return len(l)
}

func (s *Scene) SetLayer(name string, depth float32, source string, program string) {
	s.layers[name] = Layer{
		Depth:   depth,
		Texture: source,
		Program: program,
	}
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
	s.programs[name] = p

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

		return fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	gl.UseProgram(program)

	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &s.Projection[0])
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &s.Camera[0])
	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)
	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))
	s.TimeUniform = gl.GetUniformLocation(program, gl.Str("time\x00"))
	s.ModelUniform = gl.GetUniformLocation(program, gl.Str("model\x00"))
	s.DepthUniform = gl.GetUniformLocation(program, gl.Str("depth\x00"))
	return nil
}

func (s *Scene) BindBuffers() {
	// Configure the vertex data
	gl.GenVertexArrays(1, &s.vao)
	gl.BindVertexArray(s.vao)

	gl.GenBuffers(1, &s.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)

	vertAttrib := uint32(gl.GetAttribLocation(s.programs["default"].GLProgram, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(s.programs["default"].GLProgram, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	{ // previousPass pipeline setup
		gl.GenFramebuffers(1, &s.prevPassFBObj)
		gl.BindFramebuffer(gl.FRAMEBUFFER, s.prevPassFBObj)
		gl.GenTextures(1, &s.prevPassFBTex)
		gl.BindTexture(gl.TEXTURE_2D, s.prevPassFBTex)
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, s.Width, s.Height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, s.prevPassFBTex, 0)

		gl.GenRenderbuffers(1, &s.prevPassRBObj)

		gl.BindRenderbuffer(gl.RENDERBUFFER, s.prevPassRBObj)

		gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, s.Width, s.Height)
		gl.BindRenderbuffer(gl.RENDERBUFFER, 0)
		gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, s.prevPassRBObj)
	}

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

	t := float32(time.Since(tStart)) / NANOSTOSEC
	gl.Uniform1f(s.TimeUniform, t)
	gl.BindVertexArray(s.vao)

	var ls []Layer
	for _, l := range s.layers {
		ls = append(ls, l)
	}
	sort.Sort(layers(ls))

	{
		s.prevPassFBObj, s.prevFrameFBObj = s.prevFrameFBObj, s.prevPassFBObj
		s.prevPassFBTex, s.prevFrameFBTex = s.prevFrameFBTex, s.prevPassFBTex
		s.prevPassRBObj, s.prevFrameRBObj = s.prevFrameRBObj, s.prevPassRBObj

		//bind framebuffer
		gl.BindFramebuffer(gl.FRAMEBUFFER, s.prevPassFBObj)

		//snapshot framebuffer output into lastFrame
		// TODO

		//reset buffer bits
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	}

	for _, l := range ls {
		program := s.programs[l.Program].GLProgram
		gl.UseProgram(program)
		vs := verts(l.Depth)
		gl.BufferData(gl.ARRAY_BUFFER, len(vs)*4, gl.Ptr(&vs[0]), gl.STATIC_DRAW)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, s.textures[l.Texture])

		gl.ActiveTexture(gl.TEXTURE1)
		gl.BindTexture(gl.TEXTURE_2D, s.prevFrameFBTex)
		prevFrame := gl.GetUniformLocation(program, gl.Str("prevFrame"+"\x00"))
		gl.Uniform1i(prevFrame, 1)

		//texture 1, other texture
		gl.ActiveTexture(gl.TEXTURE2)
		gl.BindTexture(gl.TEXTURE_2D, s.prevPassFBTex)
		prevPass := gl.GetUniformLocation(program, gl.Str("prevPass"+"\x00"))
		gl.Uniform1i(prevPass, 2)

		gl.Uniform1f(s.DepthUniform, l.Depth)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.DrawArrays(gl.TRIANGLES, 0, 2*3)

		fmt.Println(s.prevPassFBTex, s.prevFrameFBTex, prevFrame, prevPass)
	}

	{

		gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		program := s.programs["final"].GLProgram
		gl.UseProgram(program)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, s.prevPassFBTex)

		gl.ActiveTexture(gl.TEXTURE1)
		gl.BindTexture(gl.TEXTURE_2D, s.prevFrameFBTex)
		prevFrame := gl.GetUniformLocation(program, gl.Str("prevFrame"+"\x00"))
		gl.Uniform1i(prevFrame, 1)

		//texture 1, other texture
		gl.ActiveTexture(gl.TEXTURE2)
		gl.BindTexture(gl.TEXTURE_2D, s.prevPassFBTex)
		prevPass := gl.GetUniformLocation(program, gl.Str("prevPass"+"\x00"))
		gl.Uniform1i(prevPass, 2)

		gl.ActiveTexture(gl.TEXTURE0)

		//bind framebuffer texture
		gl.BufferData(gl.ARRAY_BUFFER, len(outputTris)*4, gl.Ptr(&outputTris[0]), gl.STATIC_DRAW)
		gl.BindTexture(gl.TEXTURE_2D, s.prevPassFBTex)
		gl.DrawArrays(gl.TRIANGLES, 0, 2*3)

		//draw output wuad
	}

	s.Window.SwapBuffers()
	glfw.PollEvents()
}

func (s *Scene) ReloadPrograms() {
	for name, program := range s.programs {
		s.LoadProgram(name, program.VShaderLocation, program.FShaderLocation)
	}
}

func verts(d float32) []float32 {
	return []float32{
		-1.0, -1.0, d, 0.0, 1.0,
		1.0, -1.0, d, 1.0, 1.0,
		-1.0, 1.0, d, 0.0, 0.0,
		1.0, -1.0, d, 1.0, 1.0,
		1.0, 1.0, d, 1.0, 0.0,
		-1.0, 1.0, d, 0.0, 0.0,
	}
}

var outputTris = []float32{
	-1.0, -1.0, 0, 0.0, 0.0,
	1.0, -1.0, 0, 1.0, 0.0,
	-1.0, 1.0, 0, 0.0, 1.0,
	1.0, -1.0, 0, 1.0, 0.0,
	1.0, 1.0, 0, 1.0, 1.0,
	-1.0, 1.0, 0, 0.0, 1.0,
}
