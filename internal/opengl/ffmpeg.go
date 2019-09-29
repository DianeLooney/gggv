package opengl

import (
	"fmt"

	"github.com/dianelooney/gggv/internal/carbon"
	"github.com/dianelooney/gggv/internal/ffmpeg"
)

func init() {
	var f Source
	f = &FFVideoSource{}
	f.Name()
}

type FFVideoSource struct {
	name    SourceName
	width   int32
	height  int32
	decoder ffmpeg.Reader
	texture uint32
}

func (f *FFVideoSource) Name() SourceName {
	return f.name
}
func (f *FFVideoSource) Children() []SourceName {
	return nil
}
func (f *FFVideoSource) Render(scene *Scene) {
	/**
	if !f.decoder.Ready() {
		return
	}*/

	frame, err := f.decoder.Read()
	if err != nil {
		fmt.Println(err)
	}
	f.width, f.height = int32(frame.Width), int32(frame.Height)

	carbon.ActiveTexture(f.texture)
	carbon.BindTexture(carbon.TEXTURE_2D, f.texture)
	if len(frame.Pix) == 0 {
		return
	}
	carbon.TexImage2D(
		carbon.TEXTURE_2D,
		0,
		carbon.RGBA,
		int32(frame.Width),
		int32(frame.Height),
		0,
		carbon.RGB,
		carbon.UNSIGNED_BYTE,
		carbon.Ptr(&frame.Pix[0]))
}
func (f *FFVideoSource) SkipRender(scene *Scene) {
	//f.decoder.SkipSample()
}
func (f *FFVideoSource) Dimensions() (width, height int32) {
	return f.width, f.height
}
func (f *FFVideoSource) Texture() uint32 {
	return f.texture
}
