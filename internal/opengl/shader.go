package opengl

import (
	"github.com/go-gl/gl/all-core/gl"
)

func init() {
	var f Source
	f = &ShaderSource{}
	f.Name()
}

var ShaderSourceKind SourceKind = "ShaderSource"

const SHADER_TEXTURE_COUNT = 10

type ShaderSource struct {
	name SourceName

	program uint32
	sources [SHADER_TEXTURE_COUNT]SourceName

	texture uint32
}

func (s *ShaderSource) Kind() SourceKind {
	return ShaderSourceKind
}
func (s *ShaderSource) Name() SourceName {
	return s.name
}
func (s *ShaderSource) Children() []SourceName {
	out := make([]SourceName, SHADER_TEXTURE_COUNT)
	for i := 0; i < SHADER_TEXTURE_COUNT; i++ {
		out[i] = s.sources[i]
	}
	return out
}
func (s *ShaderSource) Render(scene *Scene) {
	gl.UseProgram(s.program)
	gl.BufferData(gl.ARRAY_BUFFER, len(staticVerts)*4, gl.Ptr(&staticVerts[0]), gl.STATIC_DRAW)

	for i, name := range s.sources {
		gl.ActiveTexture(gl.TEXTURE0 + uint32(i))
		gl.BindTexture(gl.TEXTURE_2D, scene.sources[name].Texture())
	}

	scene.bindCommonUniforms(s.program)

	if uniforms, ok := scene.uniforms[string(s.name)]; ok {
		for _, uniform := range uniforms {
			uniform.Bind(s.program)
		}
	}
	if uniforms, ok := scene.uniforms["*"]; ok {
		for _, uniform := range uniforms {
			uniform.Bind(s.program)
		}
	}

	gl.ActiveTexture(gl.TEXTURE0)
	gl.DrawArrays(gl.TRIANGLES, 0, 2*3)
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
