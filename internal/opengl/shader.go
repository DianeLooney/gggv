package opengl

import (
	"fmt"

	"github.com/dianelooney/gggv/internal/carbon"
	"github.com/dianelooney/gggv/internal/fps"
	"github.com/go-gl/mathgl/mgl32"
)

const SHADER_TEXTURE_COUNT = 10

type ShaderSource struct {
	name       SourceName
	flipOutput bool

	p        string
	sources  [SHADER_TEXTURE_COUNT]SourceName
	uniforms map[string]BindUniformer

	geometry  []float32
	drawCount int32
	width     float32
	height    float32
	fbo       uint32
	rbo       uint32
	texture   uint32
}

func (s *ShaderSource) Name() SourceName {
	return s.name
}
func (s *ShaderSource) Children() []SourceName {
	out := []SourceName{}
	for i := 0; i < SHADER_TEXTURE_COUNT; i++ {
		if s.sources[i] != "" {
			out = append(out, s.sources[i])
		}
	}
	return out
}
func (s *ShaderSource) Render(scene *Scene) {
	program := scene.programs[s.p].GLProgram
	carbon.BindFramebuffer(carbon.FRAMEBUFFER, s.fbo)
	carbon.UseProgram(program)

	carbon.ActiveTexture(carbon.TEXTURE0)
	carbon.BindTexture(carbon.TEXTURE_2D, s.texture)

	for i, name := range s.sources {
		carbon.ActiveTexture(carbon.TEXTURE1 + uint32(i))
		if name == "" || name == "-" {
			carbon.BindTexture(carbon.TEXTURE_2D, 0)
			carbon.Uniform(program, fmt.Sprintf("hasTex%v", i), 0)
			continue
		}

		source, ok := scene.sources[name]
		if !ok {
			carbon.BindTexture(carbon.TEXTURE_2D, 0)
			carbon.Uniform(program, fmt.Sprintf("hasTex%v", i), 0)
			continue
		}
		carbon.BindTexture(carbon.TEXTURE_2D, source.Texture())
		carbon.Uniform(program, fmt.Sprintf("hasTex%v", i), 1)

		switch src := source.(type) {
		case *FFVideoSource:
			carbon.Uniform(program, fmt.Sprintf("tex%vwidth", i), src.width)
			carbon.Uniform(program, fmt.Sprintf("tex%vheight", i), src.height)
		}
	}

	{
		const vertexByteSize = 6 * 4
		vertAttrib := uint32(carbon.GetAttribLocation(program, carbon.Str("vert\x00")))
		carbon.EnableVertexAttribArray(vertAttrib)
		carbon.VertexAttribPointer(vertAttrib, 3, carbon.FLOAT, false, vertexByteSize, carbon.PtrOffset(0))

		texCoordAttrib := uint32(carbon.GetAttribLocation(program, carbon.Str("vertTexCoord\x00")))
		carbon.EnableVertexAttribArray(texCoordAttrib)
		carbon.VertexAttribPointer(texCoordAttrib, 2, carbon.FLOAT, false, vertexByteSize, carbon.PtrOffset(3*4))

		particleNAttrib := uint32(carbon.GetAttribLocation(program, carbon.Str("vertParticleN\x00")))
		carbon.EnableVertexAttribArray(particleNAttrib)
		carbon.VertexAttribPointer(particleNAttrib, 1, carbon.FLOAT, false, vertexByteSize, carbon.PtrOffset(5*4))

		carbon.Uniform(program, "camera", scene.Camera)

		carbon.UniformTex(program, "lastFrame", 0)
		for i := 0; i < SHADER_TEXTURE_COUNT; i++ {
			carbon.UniformTex(program, fmt.Sprintf("tex%v", i), int32(i)+1)
		}

		carbon.Uniform(program, "time", scene.time)
		carbon.Uniform(program, "fps", float32(fps.LastSec()))
		carbon.Uniform(program, "renderTime", float32(fps.FrameDuration())/NANOSTOSEC)
		x, y := scene.Window.GetCursorPos()
		carbon.Uniform(program, "cursorX", x)
		carbon.Uniform(program, "cursorY", y)
		windowWidth, windowHeight := scene.Window.GetSize()
		carbon.Uniform(program, "windowWidth", windowWidth)
		carbon.Uniform(program, "windowHeight", windowHeight)
		carbon.Uniform(program, "windowSize", [2]float32{float32(windowWidth), float32(windowHeight)})
	}

	for _, u := range scene.uniforms {
		u.BindUniform(program)
	}

	for _, u := range s.uniforms {
		u.BindUniform(program)
	}

	if s.flipOutput {
		carbon.Uniform(program, "flipOutput", int(1))
	} else {
		carbon.Uniform(program, "flipOutput", int(0))
	}

	projectionMat := proj(s.width, s.height)
	carbon.Uniform(program, "projection", projectionMat)
	r := s.geometry

	if len(r) > 0 {
		carbon.BufferData(carbon.ARRAY_BUFFER, len(r)*4, carbon.Ptr(&r[0]), carbon.STATIC_DRAW)
	}
	carbon.DrawArraysInstanced(carbon.TRIANGLES, 0, int32(len(r)/6), s.drawCount)

	carbon.BindFramebuffer(carbon.FRAMEBUFFER, 0)
}
func (s *ShaderSource) SkipRender(scene *Scene) {}

// TODO: fix type signature here. It should be float32s and return the struct values
func (s *ShaderSource) Dimensions() (width, height int32) {
	return 1, 1
}
func (s *ShaderSource) Texture() uint32 {
	return s.texture
}

const sqrt3 = 1.732

func proj(w, h float32) mgl32.Mat4 {
	return mgl32.Ortho(-w/2, w/2, -h/2, h/2, 0.1, 10)
}

const size = 5

func rect(w, h float32) (out []float32) {
	w, h = w/2, h/2
	dx := 2 * w / size
	dy := 2 * h / size
	for x := float32(0); x < size; x++ {
		for y := float32(0); y < size; y++ {
			out = append(out, []float32{
				-w + x*dx, -h + y*dy, 0, 0 + x/size, 0 + y/size, 0,
				-w + (x+1)*dx, -h + y*dy, 0, 0 + (x+1)/size, 0 + y/size, 0,
				-w + x*dx, -h + (y+1)*dy, 0, 0 + x/size, 0 + (y+1)/size, 0,
				//
				-w + (x+1)*dx, -h + y*dy, 0, 0 + (x+1)/size, 0 + y/size, 0,
				-w + x*dx, -h + (y+1)*dy, 0, 0 + x/size, 0 + (y+1)/size, 0,
				-w + (x+1)*dx, -h + (y+1)*dy, 0, 0 + (x+1)/size, 0 + (y+1)/size, 0,
			}...)
		}
	}

	/*
		[]float32{
			-w, -h, 0, 0, 0,
			+w, -h, 0, 1, 0,
			-w, +h, 0, 0, 1,
			//
			+w, -h, 0, 1, 0,
			-w, +h, 0, 0, 1,
			+w, +h, 0, 1, 1,
		}
	*/
	return
}

var hexagon = []float32{
	-1, -1, 0, -1, -1,
	+1, -1, 0, +1, -1,
	+0, +0, 0, +0, +0,
	//
	+1, -1, 0, +1, -1,
	+2, +0, 0, +2, +0,
	+0, +0, 0, +0, +0,
	//
	+2, +0, 0, +2, +0,
	+1, +1, 0, +1, +1,
	+0, +0, 0, +0, +0,
	//
	+1, +1, 0, +1, +1,
	-1, +1, 0, -1, +1,
	+0, +0, 0, +0, +0,
	//
	-1, +1, 0, -1, +1,
	-2, +0, 0, -2, +0,
	+0, +0, 0, +0, +0,
	//
	-2, +0, 0, -2, +0,
	-1, -1, 0, -1, -1,
	+0, +0, 0, +0, +0,
}

func init() {
	for i := 0; i < 18; i++ {
		hexagon[5*i+1] *= sqrt3
		hexagon[5*i+3] = (hexagon[5*i+3] + 1) / 2
		hexagon[5*i+4] = (hexagon[5*i+4] + 1) * sqrt3 / 2
	}
}
