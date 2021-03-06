package daemon

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
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
	t := time.NewTicker(time.Second / 60)
	for !d.Scene.Window.ShouldClose() {
		<-t.C
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

func (d *D) AddShaderStorage(args net.Shifter) {
	name := args.Shift().(string)
	buff := args.Shift().(string)

	d.Schedule(func() {
		d.Scene.AddShaderStorage(name, buff)
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

func (d *D) SetShaderDimensions(args net.Shifter) {
	name := args.Shift().(string)
	_width := args.Shift()
	_height := args.Shift()
	fmt.Printf("%v %v %v\n", name, _width, _height)
	var width, height float32
	var ok bool
	if width, ok = _width.(float32); !ok {
		width = float32(_width.(int32))
	}
	if height, ok = _height.(float32); !ok {
		height = float32(_height.(int32))
	}

	d.Schedule(func() {
		d.Scene.SetShaderDimensions(name, width, height)
	})
}
func (d *D) SetShaderGeometry(args net.Shifter) {
	name := args.Shift().(string)
	geom := make([]float32, args.Length()-1)
	for i := 0; i < args.Length()-1; i++ {
		geom[i] = args.Shift().(float32)
	}
	d.Schedule(func() {
		d.Scene.SetShaderGeometry(name, geom)
	})
}

func (d *D) SetShaderDrawCount(args net.Shifter) {
	name := args.Shift().(string)
	var n int32
	var ok bool
	_n := args.Shift()
	n, ok = _n.(int32)
	if !ok {
		n = int32(math.Round(float64(_n.(float32))))
	}

	d.Schedule(func() {
		d.Scene.SetShaderDrawCount(name, n)
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

var stubsEnabled = flag.Bool("stub", true, "Enable glsl stub replacement")
var vStubPath = flag.String("vstub", "shaders/vert.stub.glsl", "Path to the vertex shader stub")
var gStubPath = flag.String("gstub", "shaders/geom.stub.glsl", "Path to the geometry shader stub")
var fStubPath = flag.String("fstub", "shaders/frag.stub.glsl", "Path to the fragment shader stub")

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

	var vs, gs, fs = string(vShader), string(gShader), string(fShader)
	if *stubsEnabled {
		vShader, err := ioutil.ReadFile(*vStubPath)
		if err != nil {
			panic(err)
		}
		gShader, err := ioutil.ReadFile(*gStubPath)
		if err != nil {
			panic(err)
		}
		fShader, err := ioutil.ReadFile(*fStubPath)
		if err != nil {
			panic(err)
		}
		vs = string(vShader) + vs
		gs = string(gShader) + gs
		fs = string(fShader) + fs
	}

	d.Schedule(func() {
		err := d.Scene.LoadProgram(name, vs, gs, fs)
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
		defer func() {
			if err := recover(); err != nil {
				logs.Error("Unable to load shader program: ", err)
			}
		}()

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
	d.WatchProgramS(name, vShaderPath, gShaderPath, fShaderPath)
}
func (d *D) WatchProgramS(name, vShaderPath, gShaderPath, fShaderPath string) {
	d.Schedule(func() {
		defer func() {
			if err := recover(); err != nil {
				logs.Error("Unable to load shader program: ", err)
			}
		}()

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
		if *stubsEnabled {
			if err := w.Add(*vStubPath); err != nil {
				logs.Log("Unable to watch shader", name, *vStubPath, err)
				return
			}
			if err := w.Add(*gStubPath); err != nil {
				logs.Log("Unable to watch shader", name, *gStubPath, err)
				return
			}
			if err := w.Add(*fStubPath); err != nil {
				logs.Log("Unable to watch shader", name, *fStubPath, err)
				return
			}
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
	if args.Length() == 6 {
		d.SetUniform4f(args)
		return
	}

	if args.Length() == 5 {
		d.SetUniform3f(args)
		return
	}

	layer := args.Shift().(string)
	uniform := args.Shift().(string)
	value := args.Shift()
	d.Schedule(func() {
		d.Scene.SetUniform(layer, uniform, value)
	})
}

func (d *D) SetGlobalUniform(args net.Shifter) {
	if args.Length() == 4 {
		d.SetGlobalUniform3f(args)
		return
	}

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

func (d *D) SetUniform4f(args net.Shifter) {
	layer := args.Shift().(string)
	uniform := args.Shift().(string)
	var f0, f1, f2, f3 float32
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
	if f2, ok = v2.(float32); !ok {
		f2 = float32(v2.(int32))
	}
	v3 := args.Shift()
	if f3, ok = v3.(float32); !ok {
		f3 = float32(v3.(int32))
	}
	d.Schedule(func() {
		d.Scene.SetUniform(layer, uniform, [4]float32{f0, f1, f2, f3})
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
	if f2, ok = v2.(float32); !ok {
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
