package opengl

import (
	"github.com/dianelooney/gvd/internal/ffmpeg"
	"github.com/go-gl/gl/all-core/gl"
)

func init() {
	var f Source
	f = &FFVideoSource{}
	f.Name()
}

var FFVideoSourceKind SourceKind = "FFVideoSourceKind"

type FFVideoSource struct {
	name    SourceName
	width   int32
	height  int32
	decoder *ffmpeg.Sampler
	texture uint32
}

func (f *FFVideoSource) Kind() SourceKind {
	return FFVideoSourceKind
}
func (f *FFVideoSource) Name() SourceName {
	return f.name
}
func (f *FFVideoSource) Children() []SourceName {
	return nil
}
func (f *FFVideoSource) Render(scene *Scene) {
	if !f.decoder.Ready() {
		return
	}

	w, h, img := f.decoder.Sample()
	f.width, f.height = int32(w), int32(h)
	gl.ActiveTexture(f.texture)
	gl.BindTexture(gl.TEXTURE_2D, f.texture)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(w),
		int32(h),
		0,
		gl.RGB,
		gl.UNSIGNED_BYTE,
		gl.Ptr(&img[0]))
}
func (f *FFVideoSource) Dimensions() (width, height int32) {
	return f.width, f.height
}
func (f *FFVideoSource) Texture() uint32 {
	return f.texture
}
