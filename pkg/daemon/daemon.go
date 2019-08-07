package daemon

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dianelooney/gvd/filters"
	"github.com/dianelooney/gvd/internal/ffmpeg"

	"github.com/dianelooney/gvd/internal/opengl"
)

func New() *Daemon {
	return &Daemon{
		Scene:    opengl.NewScene(),
		Decoders: make(map[string]*ffmpeg.AsyncDecoder),
	}
}

type Daemon struct {
	Mtx      sync.Mutex
	Scene    *opengl.Scene
	Decoders map[string]*ffmpeg.AsyncDecoder

	reloadShaders int32
}

func (d *Daemon) DrawLoop() {
	nextFrame := time.Now()
	for !d.Scene.Window.ShouldClose() {
		v := atomic.LoadInt32(&d.reloadShaders)
		if v != 0 {
			if err := d.Scene.LoadProgram("shaders/vert/default.glsl", "shaders/frag/default.glsl"); err != nil {
				fmt.Fprintf(os.Stderr, "Error while loading shaders: %v\n", err)
			}
			atomic.AddInt32(&d.reloadShaders, -v)
		}

		if nextFrame.Before(time.Now()) {
			d.Mtx.Lock()
			nextFrame = nextFrame.Add(42 * time.Millisecond)
			for name, decoder := range d.Decoders {
				img := decoder.NextFrame()
				w, h := decoder.Dimensions()
				d.filterAndBind(name, w, h, img)
			}
			d.Mtx.Unlock()
		}

		d.Scene.Draw()
	}
}

func (d *Daemon) filterAndBind(name string, width, height int, img []uint8) {
	filters := []filters.Interface{
		//invert.New(),
		//onlygreen.New(),
	}
	for _, f := range filters {
		f.Apply(img)
	}

	d.Scene.RebindTexture(name, width, height, img)
}

func (d *Daemon) ReloadShaders() {
	atomic.AddInt32(&d.reloadShaders, 1)
}
