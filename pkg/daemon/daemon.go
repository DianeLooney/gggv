package daemon

import (
	"fmt"

	"github.com/dianelooney/gggv/internal/effects"

	"github.com/dianelooney/gggv/internal/opengl"
)

func New() *D {
	return &D{
		Scene:    opengl.NewScene(),
		Schedule: make(chan effects.Interface, 10000),
	}
}

type D struct {
	Scene    *opengl.Scene
	Schedule chan effects.Interface
}

func (d *D) Begin() {
	for e := range d.Schedule {
		fmt.Printf("%T - %s\n", e, e)
		if e.Perform() == effects.ErrWindowShouldClose {
			break
		}
	}
}
func (d *D) Push(es ...effects.Interface) {
	for _, e := range es {
		d.Schedule <- e
	}
}
