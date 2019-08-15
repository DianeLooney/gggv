package opengl

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
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
	out := make([]SourceName, SHADER_TEXTURE_COUNT)
	for i := 0; i < SHADER_TEXTURE_COUNT; i++ {
		out[i] = s.sources[i]
	}
	return out
}
func (s *ShaderSource) Render(scene *Scene) {
	gl.BindFramebuffer(gl.FRAMEBUFFER, s.fbo)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(s.program)
	gl.BufferData(gl.ARRAY_BUFFER, len(staticVerts)*4, gl.Ptr(&staticVerts[0]), gl.STATIC_DRAW)

	for i, name := range s.sources {
		if name == "" {
			continue
		}
		gl.ActiveTexture(gl.TEXTURE0 + uint32(i))
		gl.BindTexture(gl.TEXTURE_2D, scene.sources[name].Texture())
	}

	scene.bindCommonUniforms(s.program)

	projectionMat := mgl32.Ortho(-1, 1, 1, -1, 0.1, 10)
	projection := gl.GetUniformLocation(s.program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projection, 1, false, &projectionMat[0])

	gl.ActiveTexture(gl.TEXTURE0)
	gl.DrawArrays(gl.TRIANGLES, 0, 2*3)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
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
