package daemon

import (
	"fmt"
	"os"
	"sync"
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

	reloadShaders chan bool
}

func (d *Daemon) DrawLoop() {
	nextFrame := time.Now()
	for !d.Scene.Window.ShouldClose() {
		select {
		case <-d.reloadShaders:
			fmt.Println("Reloading shaders")
			if err := d.Scene.LoadProgram("shaders/vert/default.glsl", "shaders/frag/default.glsl"); err != nil {
				fmt.Fprintf(os.Stderr, "Error while loading shaders: %v\n", err)
			}
		default:
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
