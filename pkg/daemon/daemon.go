package daemon

import (
	"io/ioutil"
	"sync"
	"time"

	"github.com/dianelooney/gggv/internal/logs"
	"github.com/dianelooney/gggv/internal/net"
	"github.com/radovskyb/watcher"

	"github.com/dianelooney/gggv/internal/opengl"
)

func New() *D {
	return &D{
		Scene:           opengl.NewScene(),
		watchers:        make(map[string]chan bool),
		mainThreadTasks: make(chan func(), 10000),
	}
}

type D struct {
	Mtx    sync.Mutex
	Scene  *opengl.Scene
	Paused bool

	watchers map[string]chan bool

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
		if !d.Paused {
			d.Scene.Draw()
		}
	}
}

func (d *D) Pause(args net.Shifter) {
	d.Schedule(func() {
		d.Paused = true
	})
}

func (d *D) Play(args net.Shifter) {
	d.Schedule(func() {
		d.Paused = false
	})
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

func (d *D) AddSourceFFT(args net.Shifter) {
	name := args.Shift().(string)

	d.Schedule(func() {
		d.Scene.AddSourceFFT(name)
	})
}

func (d *D) SetFFTScale(args net.Shifter) {
	name := args.Shift().(string)
	_scale := args.Shift()
	scale, ok := _scale.(float32)
	if !ok {
		scale = float32(_scale.(int32))
	}

	d.Schedule(func() {
		d.Scene.SetFFTScale(name, scale)
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
	_idx := args.Shift()
	idx, ok := _idx.(int32)
	if !ok {
		idx = int32(_idx.(float32))
	}
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

func (d *D) loadShader(name, v, g, f string) {
	vShader, err := ioutil.ReadFile(v)
	if err != nil {
		panic(err)
	}
	gShader, err := ioutil.ReadFile(g)
	if err != nil {
		panic(err)
	}
	fShader, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	d.Schedule(func() {
		err := d.Scene.LoadProgram(name, string(vShader), string(gShader), string(fShader))
		if err != nil {
			logs.Error("Unable to load program", name, err)
		}
	})
}

func (d *D) CreateProgram(args net.Shifter) {
	name := args.Shift().(string)
	vShaderPath := args.Shift().(string)
	gShaderPath := args.Shift().(string)
	fShaderPath := args.Shift().(string)
	d.Schedule(func() {
		d.loadShader(name, vShaderPath, gShaderPath, fShaderPath)
		if ch, ok := d.watchers[name]; ok {
			ch <- true
			delete(d.watchers, name)
		}
	})
}

func (d *D) WatchProgram(args net.Shifter) {
	name := args.Shift().(string)
	vShaderPath := args.Shift().(string)
	gShaderPath := args.Shift().(string)
	fShaderPath := args.Shift().(string)
	d.Schedule(func() {
		d.loadShader(name, vShaderPath, gShaderPath, fShaderPath)
		w := watcher.New()
		if err := w.Add(vShaderPath); err != nil {
			logs.Log("Unable to watch shader", name, vShaderPath, err)
			return
		}
		if err := w.Add(gShaderPath); err != nil {
			logs.Log("Unable to watch shader", name, gShaderPath, err)
			return
		}
		if err := w.Add(fShaderPath); err != nil {
			logs.Log("Unable to watch shader", name, fShaderPath, err)
			return
		}
		if ch, ok := d.watchers[name]; ok {
			ch <- true
			delete(d.watchers, name)
		}

		done := make(chan bool)
		d.watchers[name] = done

		go func() {
			for {
				select {
				case e := <-w.Event:
					logs.Log("Reloading shader", name, vShaderPath, gShaderPath, fShaderPath, e)
					d.loadShader(name, vShaderPath, gShaderPath, fShaderPath)
				case <-done:
					w.Close()
					return
				}
			}
		}()
		go w.Start(time.Second / 3)
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
	var f0, f1, f2 float32
	var ok bool
	v0 := args.Shift()
	if f0, ok = v0.(float32); !ok {
		f0 = float32(v0.(int32))
	}
	v1 := args.Shift()
	if f1, ok = v1.(float32); !ok {
		f1 = float32(v1.(int32))
	}
	v2 := args.Shift()
	if f2, ok = v0.(float32); !ok {
		f2 = float32(v2.(int32))
	}
	d.Schedule(func() {
		d.Scene.SetUniform(layer, uniform, [3]float32{f0, f1, f2})
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
