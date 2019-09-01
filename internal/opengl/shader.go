package opengl

import (
	"fmt"

	"github.com/dianelooney/gggv/internal/carbon"
	"github.com/go-gl/mathgl/mgl32"
)

var ShaderSourceKind SourceKind = "ShaderSource"

const SHADER_TEXTURE_COUNT = 10

type ShaderSource struct {
	name       SourceName
	flipOutput bool

	p        string
	sources  [SHADER_TEXTURE_COUNT]SourceName
	uniforms map[string]Uniform

	fbo     uint32
	rbo     uint32
	texture uint32
}

func (s *ShaderSource) Kind() SourceKind {
	return ShaderSourceKind
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
	carbon.Clear(carbon.COLOR_BUFFER_BIT | carbon.DEPTH_BUFFER_BIT)
	carbon.UseProgram(program)
	carbon.BufferData(carbon.ARRAY_BUFFER, len(staticVerts)*4, carbon.Ptr(&staticVerts[0]), carbon.STATIC_DRAW)

	for i, name := range s.sources {
		if name == "" {
			continue
		}

		source, ok := scene.sources[name]
		if !ok {
			continue
		}
		carbon.ActiveTexture(carbon.TEXTURE0 + uint32(i))
		carbon.BindTexture(carbon.TEXTURE_2D, source.Texture())

		switch src := source.(type) {
		case *FFVideoSource:
			carbon.Uniform(program, fmt.Sprintf("tex%vwidth", i), src.width)
			carbon.Uniform(program, fmt.Sprintf("tex%vheight", i), src.height)
		}
	}

	for _, u := range s.uniforms {
		carbon.Uniform(program, u.Name, u.Value)
	}

	scene.bindCommonUniforms(program)

	if s.flipOutput {
		projectionMat := mgl32.Ortho(-1, 1, 1, -1, 0.1, 10)
		carbon.Uniform(program, "projection", projectionMat)
	} else {
		projectionMat := mgl32.Ortho(-1, 1, -1, 1, 0.1, 10)
		carbon.Uniform(program, "projection", projectionMat)
	}

	carbon.DrawArrays(carbon.TRIANGLES, 0, 2*3)
	carbon.BindFramebuffer(carbon.FRAMEBUFFER, 0)
}
func (s *ShaderSource) Dimensions() (width, height int32) {
	return 1, 1
}
func (s *ShaderSource) Texture() uint32 {
	return s.texture
}

var staticVerts = []float32{
	-1, -1, 0, 0, 1,
	+1, -1, 0, 1, 1,
	-1, +1, 0, 0, 0,
	+1, -1, 0, 1, 1,
	+1, +1, 0, 1, 0,
	-1, +1, 0, 0, 0,
}
