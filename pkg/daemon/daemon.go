package daemon

import (
	"flag"
	"sync"

	"github.com/dianelooney/gggv/internal/opengl"
)

var showFPS = flag.Bool("fps", false, "Log fps to the command line")

func New() *D {
	return &D{
		Scene:           opengl.NewScene(),
		mainThreadTasks: make(chan func(), 10000),
	}
}

type D struct {
	Mtx   sync.Mutex
	Scene *opengl.Scene

	mainThreadTasks chan func()
}

func (d *D) Schedule(f func()) {
	d.mainThreadTasks <- f
}

func (d *D) FlushTasks() {
	for {
		select {
		case f := <-d.mainThreadTasks:
			f()
		default:
			return
		}
	}
}

func (d *D) DrawLoop() {
	for !d.Scene.Window.ShouldClose() {
		d.FlushTasks()
		d.Scene.Draw()
	}
}

func (d *D) AddSourceFFVideo(name, path string) {
	d.Schedule(func() {
		d.Scene.AddSourceFFVideo(name, path)
	})
}

func (d *D) AddSourceShader(name string) {
	d.Schedule(func() {
		d.Scene.AddSourceShader(name)
	})
}

func (d *D) SetShaderInput(name string, idx int32, target string) {
	d.Schedule(func() {
		d.Scene.SetShaderInput(name, idx, target)
	})
}

func (d *D) SetShaderProgram(name, program string) {
	d.Schedule(func() {
		d.Scene.SetShaderProgram(name, program)
	})
}

func (d *D) AddProgram(name, vShader, fShader string) {
	d.Schedule(func() {
		d.Scene.LoadProgram(name, vShader, fShader)
	})
}

func (d *D) SetSourceWrapS(name, value string) {
	d.Schedule(func() {
		d.Scene.SetSourceWrapS(name, value)
	})
}
func (d *D) SetSourceWrapT(name, value string) {
	d.Schedule(func() {
		d.Scene.SetSourceWrapT(name, value)
	})
}
func (d *D) SetSourceMinFilter(name, value string) {
	d.Schedule(func() {
		d.Scene.SetSourceMinFilter(name, value)
	})
}
func (d *D) SetSourceMagFilter(name, value string) {
	d.Schedule(func() {
		d.Scene.SetSourceMagFilter(name, value)
	})
}

func (d *D) SetUniform(layer string, name string, value interface{}) {
	d.Schedule(func() {
		d.Scene.SetUniform(layer, name, value)
	})
}
