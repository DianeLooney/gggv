package daemon

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/dianelooney/gvd/filters"
	"github.com/dianelooney/gvd/internal/ffmpeg"
	"github.com/dianelooney/gvd/internal/fps"

	"github.com/dianelooney/gvd/internal/opengl"
)

var showFPS = flag.Bool("fps", false, "Log fps to the command line")

func New() *D {
	return &D{
		Scene:           opengl.NewScene(),
		decoders:        make(map[string]*ffmpeg.AsyncDecoder),
		nextFrames:      make(map[string]time.Time),
		mainThreadTasks: make(chan func(), 20),
	}
}

type D struct {
	Mtx        sync.Mutex
	Scene      *opengl.Scene
	nextFrames map[string]time.Time
	decoders   map[string]*ffmpeg.AsyncDecoder

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

		d.Mtx.Lock()
		for name, decoder := range d.decoders {
			nextFrame := d.nextFrames[name]
			if !nextFrame.Before(time.Now()) {
				continue
			}

			img, duration := decoder.NextFrame()
			d.nextFrames[name] = nextFrame.Add(duration)
			w, h := decoder.Dimensions()
			d.filterAndBind(name, w, h, img)
		}

		d.Mtx.Unlock()

		d.Scene.Draw()
		fps.Next()
		if *showFPS {
			fmt.Printf("FPS: %v\t%v\n", fps.LastSec(), fps.FrameDuration())
		}
	}
}

func (d *D) filterAndBind(name string, width, height int, img []uint8) {
	filters := []filters.Interface{
		//invert.New(),
		//onlygreen.New(),
	}
	for _, f := range filters {
		f.Apply(img)
	}

	if len(img) == 0 {
		return
	}

	d.Scene.RebindTexture(name, width, height, img)
}

func (d *D) AddSource(name, path string) {
	d.Schedule(func() {
		d.Scene.TextureInit(name)
		if _, ok := d.nextFrames[name]; !ok {
			d.nextFrames[name] = time.Now()
		}

		if dec, ok := d.decoders[name]; ok {
			dec.Dealloc()
		}
		dec, err := ffmpeg.NewAsyncFileDecoder(path)
		if err != nil {
			fmt.Println("Error adding source:", err)
			return
		}
		d.decoders[name] = dec
	})
}

func (d *D) AddLayer(name, source, program string, depth float32) {
	d.Schedule(func() {
		d.Scene.SetLayer(name, depth, source)
	})
}

func (d *D) AddProgram(name, pathV, pathF string) {
	d.Schedule(func() {
		d.Scene.LoadProgram(name, pathV, pathF)
	})
}

func (d *D) ReloadPrograms() {
	d.Schedule(func() {
		d.Scene.ReloadPrograms()
	})
}
