package main

import (
	"flag"
	_ "image/png"
	"runtime"

	"github.com/hypebeast/go-osc/osc"

	"github.com/dianelooney/gggv/pkg/daemon"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var dmn *daemon.D

func main() {
	flag.Parse()
	go netSetup()

	dmn = daemon.New()

	if err := dmn.Scene.LoadProgram("default", "shaders/vert/default.glsl", "shaders/frag/default.glsl"); err != nil {
		panic(err)
	}

	if err := dmn.Scene.LoadProgram("final", "shaders/vert/final.glsl", "shaders/frag/final.glsl"); err != nil {
		panic(err)
	}

	dmn.Scene.BindBuffers()
	dmn.AddSourceFFVideo("default0", "sample0.mp4")
	dmn.AddSourceFFVideo("default1", "sample1.mp4")
	dmn.AddSourceFFVideo("default2", "sample2.mp4")
	dmn.AddSourceShader("window")

	dmn.SetUniform("ampl", "float", float32(0.0), []string{"default"})

	dmn.DrawLoop()
}

var netAddr = flag.String("net", ":4200", "Network address to listen at.")

func netSetup() {
	server := &osc.Server{Addr: *netAddr}

	server.Handle("/source/add/ffvideo", func(msg *osc.Message) {
		dmn.AddSourceFFVideo(msg.Arguments[0].(string), msg.Arguments[1].(string))
	})
	server.Handle("/source/add/shader", func(msg *osc.Message) {
		dmn.AddSourceShader(msg.Arguments[0].(string))
	})
	server.Handle("/programs/reload", func(msg *osc.Message) {
		dmn.ReloadPrograms()
	})

	server.ListenAndServe()
}
