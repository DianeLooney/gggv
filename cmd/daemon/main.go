package main

import (
	_ "image/png"
	"log"
	"runtime"

	"github.com/dianelooney/gvd/internal/ffmpeg"
	"github.com/dianelooney/gvd/pkg/daemon"
	"github.com/fsnotify/fsnotify"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var dmn *daemon.Daemon

func main() {
	dmn = daemon.New()

	dmn.Decoders["default"] = ffmpeg.NewAsyncFileDecoder("sample.mp4")
	dmn.Decoders["swap"] = ffmpeg.NewAsyncFileDecoder("sample2.mp4")
	dmn.Decoders["swap2"] = ffmpeg.NewAsyncFileDecoder("sample2.mp4")
	dmn.Decoders["swap3"] = ffmpeg.NewAsyncFileDecoder("sample2.mp4")
	dmn.Decoders["swap4"] = ffmpeg.NewAsyncFileDecoder("sample2.mp4")

	if err := dmn.Scene.LoadProgram("shaders/vert/default.glsl", "shaders/frag/default.glsl"); err != nil {
		panic(err)
	}

	dmn.Scene.BindBuffers()
	dmn.Scene.TextureInit("default")
	dmn.Scene.TextureInit("swap")
	dmn.Scene.TextureInit("swap2")
	dmn.Scene.TextureInit("swap3")
	dmn.Scene.TextureInit("swap4")

	go watchShaders()

	dmn.DrawLoop()
}

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
			dmn.ReloadShaders()
		}
	}()
}
