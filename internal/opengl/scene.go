package opengl

import (
	"log"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

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

	s := &Scene{}

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

	return s
}

func finalizeScene(s *Scene) {
	defer glfw.Terminate()
}

type Scene struct {
	Window *glfw.Window

	Camera     mgl32.Mat4
	Model      mgl32.Mat4
	Projection mgl32.Mat4
}
