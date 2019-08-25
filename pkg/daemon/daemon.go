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
		mainThreadTasks: make(chan func(), 20),
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

func (d *D) AddProgram(name, pathV, pathF string) {
	d.Schedule(func() {
		d.Scene.LoadProgram(name, pathV, pathF)
	})
}

func (d *D) SetUniform(name string, value interface{}, layers []string) {
	for _, l := range layers {
		d.Scene.SetUniform(l, name, value)
	}
}

func (d *D) ReloadPrograms() {
	d.Schedule(func() {
		d.Scene.ReloadPrograms()
	})
}
