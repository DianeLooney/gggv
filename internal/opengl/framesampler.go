package opengl

import (
	"unsafe"

	"github.com/go-gl/gl/all-core/gl"
)

func NewSampler() *Sampler {
	return &Sampler{}
}

type Sampler struct {
	width  int
	height int
	pix    []uint8
}

func (s *Sampler) Ready() bool {
	return true
}

func (s *Sampler) Done() {}

func (s *Sampler) Sample() (width, height int, pix []uint8) {
	gl.ReadPixels(
		0,
		0,
		int32(s.width),
		int32(s.height),
		gl.RGB,
		gl.BYTE,
		unsafe.Pointer(&pix[0]),
	)
	return s.width, s.height, pix
}
