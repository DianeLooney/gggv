package main

import (
	"fmt"
	_ "image/png"
	"log"
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/dianelooney/gvd/internal/ffmpeg"
	"github.com/dianelooney/gvd/internal/opengl"
	"github.com/fsnotify/fsnotify"

	"github.com/giorgisio/goav/avutil"
	"github.com/go-gl/gl/all-core/gl"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var scene *opengl.Scene
var decoder *ffmpeg.Decoder
var frame *avutil.Frame

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a movie file")
		os.Exit(1)
	}
	decoder, frame = ffmpeg.NewFileDecoder(os.Args[1])
	scene = opengl.NewScene()

	// Configure the vertex and fragment shaders
	err := scene.LoadProgram("shaders/vert/default.glsl", "shaders/frag/default.glsl")
	if err != nil {
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
