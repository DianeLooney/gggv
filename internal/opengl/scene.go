package opengl

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const NANOSTOSEC = 1000000000

var tStart = time.Now()

const windowWidth = 800
const windowHeight = 600

func NewScene() *Scene {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	s := &Scene{
		programs: make(map[string]uint32),
		textures: make(map[string]uint32),
	}

	var err error
	s.Window, err = glfw.CreateWindow(windowWidth, windowHeight, "Cube", nil, nil)
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
	s.Model = mgl32.Ident4()

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	return s
}

func finalizeScene(s *Scene) {
	defer glfw.Terminate()
}

type Scene struct {
	Window *glfw.Window

	vao uint32

	Camera       mgl32.Mat4
	Model        mgl32.Mat4
	Projection   mgl32.Mat4
	ModelUniform int32
	TimeUniform  int32

	programs map[string]uint32
	textures map[string]uint32
}

func (s *Scene) LoadProgram(name, vertShaderLocation, fragShaderFlocation string) error {
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
		return err
	}

	program := gl.CreateProgram()
	s.programs[name] = program

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
	s.ModelUniform = gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(s.ModelUniform, 1, false, &s.Model[0])
	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)
	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))
	s.TimeUniform = gl.GetUniformLocation(program, gl.Str("time\x00"))

	return nil
}

func (s *Scene) BindBuffers() {
	// Configure the vertex data
	gl.GenVertexArrays(1, &s.vao)
	gl.BindVertexArray(s.vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(s.programs["default"], gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(s.programs["default"], gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
}

func (s *Scene) TextureInit(name string) {
	var t uint32
	gl.GenTextures(1, &t)
	gl.ActiveTexture(gl.TEXTURE0)
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
	gl.ActiveTexture(gl.TEXTURE0)
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
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UseProgram(s.programs["default"])
	gl.UniformMatrix4fv(s.ModelUniform, 1, false, &s.Model[0])

	t := float32(time.Since(tStart)) / NANOSTOSEC
	gl.Uniform1f(s.TimeUniform, t)
	gl.BindVertexArray(s.vao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, s.textures["default"])

	gl.DrawArrays(gl.TRIANGLES, 0, 1*2*3)

	// Maintenance
	s.Window.SwapBuffers()
	glfw.PollEvents()
}

var cubeVertices = []float32{
	// Front
	-1.0, -1.0, 1.0, 0.0, 1.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, 1.0, 0.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, 1.0, 1.0, 1.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 0.0,
}
