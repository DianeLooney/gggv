package com

type Sampler interface {
	Ready() bool
	Done()
	Sample() (width, height int, pix []uint8)
}
