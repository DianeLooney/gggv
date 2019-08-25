package opengl

import (
	"fmt"

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

	p       string
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
	gl.BindFramebuffer(gl.FRAMEBUFFER, s.fbo)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
	gl.BufferData(gl.ARRAY_BUFFER, len(staticVerts)*4, gl.Ptr(&staticVerts[0]), gl.STATIC_DRAW)

	fmt.Println(s.sources)
	for i, name := range s.sources {
		if name == "" {
			continue
		}

		fmt.Printf("Bound source %v\n", name)
		source := scene.sources[name]
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

	scene.bindCommonUniforms(program)

	projectionMat := mgl32.Ortho(-1, 1, 1, -1, 0.1, 10)
	projection := gl.GetUniformLocation(program, gl.Str("projection\x00"))
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
