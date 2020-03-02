package opengl

import (
	"math/cmplx"

	"github.com/dianelooney/gggv/internal/carbon"

	"github.com/dianelooney/gggv/internal/logs"

	"github.com/gordonklaus/portaudio"
	"github.com/mjibson/go-dsp/fft"
)

const sampleCount = 1470

func NewPortaudio() Source {
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
	src.init()
	go src.main()
	return &src
}

type Portaudio struct {
	N SourceName

	scale    float32
	stream   *portaudio.Stream
	samplesI []int32
	samples  []float64
	fft      []float32
	texture  uint32
}

func (s *Portaudio) init() {
	portaudio.Initialize()
	var err error
	s.stream, err = portaudio.OpenDefaultStream(1, 0, 44100, len(s.samplesI), s.samplesI)
	chk(err)
	chk(s.stream.Start())
}

func (s *Portaudio) main() {
	defer portaudio.Terminate()
	defer s.stream.Close()

	for {
		if err := s.stream.Read(); err != nil {
			logs.Error(err)
		}
		for i, x := range s.samplesI {
			s.samples[i] = float64(x) / float64(1<<31)
		}
		out := fft.FFTReal(s.samples)
		for i, z := range out {
			s.fft[i] = float32(cmplx.Abs(z)) / s.scale
		}
	}
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
	carbon.ActiveTexture(s.texture)
	carbon.BindTexture(carbon.TEXTURE_2D, s.texture)
	carbon.TexImage2D(
		carbon.TEXTURE_2D,
		0,
		carbon.RGBA,
		int32(len(s.fft)/(3*2)),
		1,
		0,
		carbon.LUMINANCE,
		carbon.FLOAT,
		carbon.Ptr(&s.fft[0]),
	)
	carbon.Flush()
}
func (s *Portaudio) Dimensions() (width, height int32) {
	return int32(len(s.fft)), 1
}
func (s *Portaudio) Texture() uint32 {
	return s.texture
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
