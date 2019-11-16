package daemon

import (
	"sync"
	"time"

	"github.com/dianelooney/gggv/internal/logs"
	"github.com/dianelooney/gggv/internal/net"

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
func (d *D) AddSourceFFVideo(args net.Shifter) {
	name := args.Shift().(string)
	path := args.Shift().(string)

	d.Schedule(func() {
		d.Scene.AddSourceFFVideo(name, path)
	})
}

// AddSourceShader - /source.shader/create
//
// Accepts one argument
// 1. Name (string) - the name of the shader to create
func (d *D) AddSourceShader(args net.Shifter) {
	name := args.Shift().(string)
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
func (d *D) SetShaderInput(args net.Shifter) {
	name := args.Shift().(string)
	idx := args.Shift().(int32)
	shader := args.Shift().(string)
	d.Schedule(func() {
		d.Scene.SetShaderInput(name, idx, shader)
	})
}

func (d *D) SetShaderProgram(args net.Shifter) {
	name := args.Shift().(string)
	program := args.Shift().(string)
	d.Schedule(func() {
		d.Scene.SetShaderProgram(name, program)
	})
}

func (d *D) SetFFVideoTimescale(args net.Shifter) {
	name := args.Shift().(string)
	timescale := args.Shift().(float32)
	d.Schedule(func() {
		d.Scene.SetFFVideoTimescale(name, float64(timescale))
	})
}

func (d *D) AddProgram(args net.Shifter) {
	name := args.Shift().(string)
	vShader := args.Shift().(string)
	gShader := args.Shift().(string)
	fShader := args.Shift().(string)
	d.Schedule(func() {
		err := d.Scene.LoadProgram(name, vShader, gShader, fShader)
		if err != nil {
			logs.Error("Unable to load program", name, err)
		}
	})
}

func (d *D) SetSourceWrapS(args net.Shifter) {
	name := args.Shift().(string)
	value := args.Shift().(string)
	d.Schedule(func() {
		d.Scene.SetSourceWrapS(name, value)
	})
}

func (d *D) SetSourceWrapT(args net.Shifter) {
	name := args.Shift().(string)
	value := args.Shift().(string)
	d.Schedule(func() {
		d.Scene.SetSourceWrapT(name, value)
	})
}

func (d *D) SetSourceMinFilter(args net.Shifter) {
	name := args.Shift().(string)
	value := args.Shift().(string)
	d.Schedule(func() {
		d.Scene.SetSourceMinFilter(name, value)
	})
}

func (d *D) SetSourceMagFilter(args net.Shifter) {
	name := args.Shift().(string)
	value := args.Shift().(string)
	d.Schedule(func() {
		d.Scene.SetSourceMagFilter(name, value)
	})
}

func (d *D) SetUniform(args net.Shifter) {
	layer := args.Shift().(string)
	uniform := args.Shift().(string)
	value := args.Shift()
	d.Schedule(func() {
		d.Scene.SetUniform(layer, uniform, value)
	})
}

func (d *D) SetGlobalUniform(args net.Shifter) {
	name := args.Shift().(string)
	value := args.Shift()
	d.Schedule(func() {
		d.Scene.SetGlobalUniform(name, value)
	})
}

func (d *D) SetUniformTimestamp(args net.Shifter) {
	layer := args.Shift().(string)
	uniform := args.Shift().(string)
	d.Schedule(func() {
		d.Scene.SetUniform(layer, uniform, time.Now())
	})
}

func (d *D) SetGlobalUniformTimestamp(args net.Shifter) {
	uniform := args.Shift().(string)
	d.Schedule(func() {
		d.Scene.SetGlobalUniform(uniform, time.Now())
	})
}

func (d *D) SetUniformClock(args net.Shifter) {
	layer := args.Shift().(string)
	uniform := args.Shift().(string)
	d.Schedule(func() {
		d.Scene.SetUniformClock(layer, uniform, time.Now())
	})
}

func (d *D) SetGlobalUniformClock(args net.Shifter) {
	uniform := args.Shift().(string)
	d.Schedule(func() {
		d.Scene.SetGlobalUniformClock(uniform, time.Now())
	})
}

func (d *D) SetUniform3f(args net.Shifter) {
	layer := args.Shift().(string)
	uniform := args.Shift().(string)
	v0 := args.Shift().(float32)
	v1 := args.Shift().(float32)
	v2 := args.Shift().(float32)
	d.Schedule(func() {
		d.Scene.SetUniform(layer, uniform, [3]float32{v0, v1, v2})
	})
}

func (d *D) SetGlobalUniform3f(args net.Shifter) {
	uniform := args.Shift().(string)
	v0 := args.Shift().(float32)
	v1 := args.Shift().(float32)
	v2 := args.Shift().(float32)
	d.Schedule(func() {
		d.Scene.SetGlobalUniform(uniform, [3]float32{v0, v1, v2})
	})
}
