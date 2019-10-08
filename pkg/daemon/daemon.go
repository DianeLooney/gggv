package daemon

import (
	"sync"
	"time"

	"github.com/dianelooney/gggv/internal/logs"

	"github.com/dianelooney/gggv/internal/opengl"
)

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

// Schedule is for internal use
func (d *D) Schedule(f func()) {
	d.mainThreadTasks <- f
}

// FlushTasks is for internal use
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

// DrawLoop is for internal use
func (d *D) DrawLoop() {
	for !d.Scene.Window.ShouldClose() {
		d.FlushTasks()
		d.Scene.Draw()
	}
}

// AddSourceFFVideo - /source.ffvideo/create
//
// Accepts two arguments
// 1. Name (string) - the name of the source to create
// 2. Path (string) - the path of the video
func (d *D) AddSourceFFVideo(name, path string) {
	d.Schedule(func() {
		d.Scene.AddSourceFFVideo(name, path)
	})
}

// AddSourceShader - /source.shader/create
//
// Accepts one argument
// 1. Name (string) - the name of the shader to create
func (d *D) AddSourceShader(name string) {
	d.Schedule(func() {
		d.Scene.AddSourceShader(name)
	})
}

// SetShaderInput - /source.shader/set/input
//
// Accepts three arguments
// 1. Name (string) - the name of the shader
// 2. Index (int32) - the index the source should be bound to
// 3. Source (string) - the name of the source to be bound to the shader
func (d *D) SetShaderInput(name string, idx int32, source string) {
	d.Schedule(func() {
		d.Scene.SetShaderInput(name, idx, source)
	})
}

func (d *D) SetShaderProgram(name, program string) {
	d.Schedule(func() {
		d.Scene.SetShaderProgram(name, program)
	})
}

func (d *D) SetFFVideoTimescale(name string, timescale float32) {
	d.Schedule(func() {
		d.Scene.SetFFVideoTimescale(name, float64(timescale))
	})
}

func (d *D) AddProgram(name, vShader, gShader, fShader string) {
	d.Schedule(func() {
		err := d.Scene.LoadProgram(name, vShader, gShader, fShader)
		if err != nil {
			logs.Error("Unable to load program", name, err)
		}
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

func (d *D) SetUniformTimestamp(layer string, name string) {
	d.Schedule(func() {
		d.Scene.SetUniform(layer, name, time.Now())
	})
}

func (d *D) SetUniformClock(layer string, name string) {
	d.Schedule(func() {
		d.Scene.SetUniformClock(layer, name, time.Now())
	})
}

func (d *D) SetUniform3f(layer string, name string, v0, v1, v2 float32) {
	d.Schedule(func() {
		d.Scene.SetUniform(layer, name, [3]float32{v0, v1, v2})
	})
}
