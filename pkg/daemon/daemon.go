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

func New() *D {
	return &D{
		Scene:    opengl.NewScene(),
		decoders: make(map[string]*ffmpeg.AsyncDecoder),
	}
}

type D struct {
	Mtx      sync.Mutex
	Scene    *opengl.Scene
	decoders map[string]*ffmpeg.AsyncDecoder

	reloadShaders int32
}

func (d *D) DrawLoop() {
	nextFrame := time.Now()
	for !d.Scene.Window.ShouldClose() {
		v := atomic.LoadInt32(&d.reloadShaders)
		if v != 0 {
			if err := d.Scene.LoadProgram("default", "shaders/vert/default.glsl", "shaders/frag/default.glsl"); err != nil {
				fmt.Fprintf(os.Stderr, "Error while loading shaders: %v\n", err)
			}
			atomic.AddInt32(&d.reloadShaders, -v)
		}

		if nextFrame.Before(time.Now()) {
			d.Mtx.Lock()
			nextFrame = nextFrame.Add(42 * time.Millisecond)
			for name, decoder := range d.decoders {
				img := decoder.NextFrame()
				w, h := decoder.Dimensions()
				d.filterAndBind(name, w, h, img)
			}
			d.Mtx.Unlock()
		}

		d.Scene.Draw()
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

	d.Scene.RebindTexture(name, width, height, img)
}

func (d *D) ReloadShaders() {
	atomic.AddInt32(&d.reloadShaders, 1)
}

func (d *D) AddSource(name, path string) (err error) {
	d.Mtx.Lock()
	defer d.Mtx.Unlock()

	if dec, ok := d.decoders[name]; ok {
		dec.Dealloc()
	}
	dec, err := ffmpeg.NewAsyncFileDecoder(path)
	if err != nil {
		return err
	}
	d.decoders[name] = dec
	return nil
}
