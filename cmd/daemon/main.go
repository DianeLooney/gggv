package main

import (
	"fmt"
	_ "image/png"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/dianelooney/gvd/filters"
	"github.com/dianelooney/gvd/filters/invert"
	"github.com/dianelooney/gvd/filters/onlygreen"
	"github.com/dianelooney/gvd/internal/ffmpeg"
	"github.com/dianelooney/gvd/internal/opengl"
	"github.com/fsnotify/fsnotify"

	"github.com/giorgisio/goav/avutil"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var scene *opengl.Scene
var mtx sync.Mutex
var decoder1 *ffmpeg.Decoder
var decoder2 *ffmpeg.Decoder
var frame *avutil.Frame

func main() {
	decoder1, _ = ffmpeg.NewFileDecoder("sample.mp4")

	go coordinatePlaylist()
	scene = opengl.NewScene()

	if err := scene.LoadProgram("shaders/vert/default.glsl", "shaders/frag/default.glsl"); err != nil {
		panic(err)
	}

	scene.BindBuffers()
	scene.TextureInit("default")
	{
		img := decoder1.NextFrame()
		width, height := decoder1.Dimensions()
		filterAndBind("default", width, height, img)
	}

	go watchShaders()

	for !scene.Window.ShouldClose() {
		select {
		case <-reloadShaders:
			fmt.Println("Reloading shaders")
			if err := scene.LoadProgram("shaders/vert/default.glsl", "shaders/frag/default.glsl"); err != nil {
				fmt.Fprintf(os.Stderr, "Error while loading shaders: %v\n", err)
			}
		default:
		}
		if nextFrame.Before(time.Now()) {
			mtx.Lock()
			var img []uint8
			for nextFrame.Before(time.Now()) {
				img = decoder1.NextFrame()
				nextFrame = nextFrame.Add(42 * time.Millisecond)
			}
			width, height := decoder1.Dimensions()
			filterAndBind("default", width, height, img)
			mtx.Unlock()
		}

		scene.Draw()
	}
}

var reloadShaders = make(chan bool)
var nextFrame = time.Now()

func watchShaders() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Panic(err)
	}
	watcher.Add("shaders/frag/default.glsl")
	watcher.Add("shaders/vert/default.glsl")

	go func() {
		for {
			if (<-watcher.Events).Op != fsnotify.Write {
				continue
			}
			reloadShaders <- true
		}
	}()
}

func coordinatePlaylist() {
	for {
		time.Sleep(time.Second)
		mtx.Lock()

		mtx.Unlock()
	}
}

func filterAndBind(name string, width, height int, img []uint8) {
	filters := []filters.Interface{
		invert.New(),
		onlygreen.New(),
	}
	for _, f := range filters {
		f.Apply(img)
	}

	scene.RebindTexture(name, width, height, img)
}
