package opengl

import (
	"github.com/dianelooney/gggv/internal/carbon"
	"github.com/dianelooney/gggv/internal/ffmpeg"
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
	carbon.ActiveTexture(f.texture)
	carbon.BindTexture(carbon.TEXTURE_2D, f.texture)
	carbon.TexImage2D(
		carbon.TEXTURE_2D,
		0,
		carbon.RGBA,
		int32(w),
		int32(h),
		0,
		carbon.RGB,
		carbon.UNSIGNED_BYTE,
		carbon.Ptr(&img[0]))
}
func (f *FFVideoSource) SkipRender(scene *Scene) {
	f.decoder.SkipSample()
}
func (f *FFVideoSource) Dimensions() (width, height int32) {
	return f.width, f.height
}
func (f *FFVideoSource) Texture() uint32 {
	return f.texture
}
