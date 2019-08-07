package invert

import "github.com/dianelooney/gvd/filters"

func New() filters.Interface {
	return &filter{}
}

type filter struct{}

func (*filter) Apply(img []uint8) {
	for i := range img {
		img[i] = 255 - img[i]
	}
}
