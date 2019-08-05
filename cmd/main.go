package main

import (
	"fmt"
	_ "image/png"
	"io/ioutil"
	"log"
	"runtime"
	"time"
	"unsafe"

	"github.com/dianelooney/gvd/internal/ffmpeg"
	"github.com/dianelooney/gvd/internal/opengl"
	"github.com/dianelooney/gvd/internal/playlist"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v2"

	"github.com/giorgisio/goav/avutil"
	"github.com/go-gl/gl/all-core/gl"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var pl *playlist.Playlist
var scene *opengl.Scene
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

	decoder.NextFrame()
	newTexture()

	go watchShaders()

	for !scene.Window.ShouldClose() {
		select {
		case <-reloadShaders:
			fmt.Println("Reloading shaders")
			err := scene.LoadProgram("shaders/vert/default.glsl", "shaders/frag/default.glsl")
			if err != nil {
				panic(err)
			}

		default:
		}
		if nextFrame.Before(time.Now()) {
			for nextFrame.Before(time.Now()) {
				decoder.NextFrame()
				nextFrame = nextFrame.Add(42 * time.Millisecond)
			}
			newTexture()
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
			decoder, frame = ffmpeg.NewFileDecoder(pv.Path)
			time.Sleep(time.Duration(pv.Duration * float64(time.Second)))
		}
	}
}

func newTexture() {
	width, height := decoder.Dimensions()

	var img []uint8
	for y := 0; y < height; y++ {
		data0 := avutil.Data(frame)[0]
		buf := make([]byte, width*3)
		startPos := uintptr(unsafe.Pointer(data0)) + uintptr(y)*uintptr(avutil.Linesize(frame)[0])
		for i := 0; i < width*3; i++ {
			element := *(*uint8)(unsafe.Pointer(startPos + uintptr(i)))
			buf[i] = element
		}
		img = append(img, buf...)
	}

	gl.GenTextures(1, &scene.Texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, scene.Texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(width),
		int32(height),
		0,
		gl.RGB,
		gl.UNSIGNED_BYTE,
		gl.Ptr(&img[0]))
}
