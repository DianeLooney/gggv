package main

import (
	"fmt"
	_ "image/png"
	"io/ioutil"
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
	"github.com/dianelooney/gvd/internal/playlist"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v2"

	"github.com/giorgisio/goav/avutil"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var pl *playlist.Playlist
var scene *opengl.Scene
var mtx sync.Mutex
var decoder *ffmpeg.Decoder
var frame *avutil.Frame

func main() {
	data, err := ioutil.ReadFile("playlist.yml")
	if err != nil {
		panic(err)
	}
	pl = &playlist.Playlist{}
	if err := yaml.Unmarshal(data, pl); err != nil {
		panic(err)
	}
	decoder, frame = ffmpeg.NewFileDecoder(pl.Videos[0].Path)
	go coordinatePlaylist()
	scene = opengl.NewScene()

	if err := scene.LoadProgram("shaders/vert/default.glsl", "shaders/frag/default.glsl"); err != nil {
		panic(err)
	}

	scene.BindBuffers()
	scene.TextureInit("default")
	{
		img := decoder.NextFrame()
		width, height := decoder.Dimensions()
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
				img = decoder.NextFrame()
				nextFrame = nextFrame.Add(42 * time.Millisecond)
			}
			width, height := decoder.Dimensions()
			filterAndBind("default", width, height, img)
			mtx.Unlock()
		}
		if err != nil {
			log.Fatalln(err)
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
		for _, pv := range pl.Videos {
			fmt.Println("Loading file", pv.Path, pv.Duration)
			dealloc := decoder
			mtx.Lock()
			decoder, frame = ffmpeg.NewFileDecoder(pv.Path)
			mtx.Unlock()
			if dealloc != nil {
				dealloc.Dealloc()
			}
			time.Sleep(time.Duration(pv.Duration * float64(time.Second)))
		}
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
