package opengl

import (
	"github.com/dianelooney/gggv/internal/carbon"
	"github.com/dianelooney/gggv/internal/logs"
)

const sampleCount = 1470

func NewPortaudio() Source {
	logs.Warn("FFT Sources not supported on Windows")

	src := Portaudio{
		samplesI: make([]int32, sampleCount),
		samples:  make([]float64, sampleCount),
		fft:      make([]float32, sampleCount),
	}
	carbon.GenTextures(1, &src.texture)
	carbon.ActiveTexture(src.texture)
	carbon.BindTexture(carbon.TEXTURE_2D, src.texture)
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MIN_FILTER, carbon.LINEAR)
	carbon.TexParameteri(carbon.TEXTURE_2D, carbon.TEXTURE_MAG_FILTER, carbon.LINEAR)
	return &src
}

type Portaudio struct {
	N       string
	texture uint32
}

func (p *Portaudio) Name() SourceName {
	return p.N
}

func (*Portaudio) SkipRender(scene *Scene) {
	return
}

func (*Portaudio) Children() []SourceName {
	return nil
}

func (s *Portaudio) Render(scene *Scene) {
}
func (s *Portaudio) Dimensions() (width, height int32) {
	return 1, 1
}
func (s *Portaudio) Texture() uint32 {
	return s.texture
}
