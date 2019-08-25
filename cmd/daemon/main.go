package main

import (
	"flag"
	"fmt"
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
	dmn.AddSourceShader("default")
	dmn.SetShaderInput("default", 0, "default0")
	dmn.SetShaderInput("default", 1, "default1")
	dmn.SetShaderInput("default", 2, "default2")
	dmn.AddSourceShader("window")
	dmn.SetShaderInput("window", 0, "default")
	dmn.SetUniform("ampl", float32(0.0), []string{"default"})

	go netSetup()
	dmn.DrawLoop()
}

var netAddr = flag.String("net", ":4200", "Network address to listen at.")

func netSetup() {
	server := &osc.Server{Addr: *netAddr}

	server.Handle("/source.ffvideo/add/ffvideo", func(msg *osc.Message) {
		dmn.AddSourceFFVideo(msg.Arguments[0].(string), msg.Arguments[1].(string))
	})
	server.Handle("/source.shader/add", func(msg *osc.Message) {
		dmn.AddSourceShader(msg.Arguments[0].(string))
	})
	server.Handle("/source.shader/set/source", func(msg *osc.Message) {
		layer := msg.Arguments[0].(string)
		index := msg.Arguments[1].(int32)
		value := msg.Arguments[2].(string)

		fmt.Println("/source.shader/set/source", layer, index, value)
		dmn.SetShaderInput(layer, index, value)
	})
	server.Handle("/source.shader/set/uniform1f", func(msg *osc.Message) {
		layer := msg.Arguments[0].(string)
		name := msg.Arguments[1].(string)
		target := msg.Arguments[2].(float32)

		dmn.Scene.SetUniform(layer, name, target)
	})
	server.Handle("/programs/reload", func(msg *osc.Message) {
		dmn.ReloadPrograms()
	})

	server.ListenAndServe()
}
