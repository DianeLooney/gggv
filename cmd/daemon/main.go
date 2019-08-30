package main

import (
	"flag"
	_ "image/png"
	"io/ioutil"
	"runtime"

	"github.com/dianelooney/gggv/internal/logs"

	"github.com/hypebeast/go-osc/osc"

	_ "github.com/dianelooney/gggv/internal/carbon/opengl" // necessary to fill carbon stubs
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

	{
		name := "default"
		vShaderPath := "shaders/vert/default.glsl"
		fShaderPath := "shaders/frag/default.glsl"
		vShader, err := ioutil.ReadFile(vShaderPath)
		if err != nil {
			return
		}
		fShader, err := ioutil.ReadFile(fShaderPath)
		if err != nil {
			return
		}
		dmn.AddProgram(name, string(vShader), string(fShader))
	}

	{
		name := "window"
		vShaderPath := "shaders/vert/window.glsl"
		fShaderPath := "shaders/frag/window.glsl"
		vShader, err := ioutil.ReadFile(vShaderPath)
		if err != nil {
			return
		}
		fShader, err := ioutil.ReadFile(fShaderPath)
		if err != nil {
			return
		}
		dmn.AddProgram(name, string(vShader), string(fShader))
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
	dmn.SetUniform("default", "ampl", float32(0.0))

	go netSetup()
	dmn.DrawLoop()
}

var netAddr = flag.String("net", ":4200", "Network address to listen at.")

func netSetup() {
	server := &osc.Server{Addr: *netAddr}

	server.Handle("/source.ffvideo/create", func(msg *osc.Message) {
		sourceName := msg.Arguments[0].(string)
		path := msg.Arguments[0].(string)

		logs.Log("/source.ffvideo/create", sourceName, path)
		dmn.AddSourceFFVideo(sourceName, path)
	})
	server.Handle("/source.shader/create", func(msg *osc.Message) {
		name := msg.Arguments[0].(string)
		logs.Log("/source.shader/create", name)
		dmn.AddSourceShader(name)
	})
	server.Handle("/source.shader/set/input", func(msg *osc.Message) {
		layer := msg.Arguments[0].(string)
		index := msg.Arguments[1].(int32)
		value := msg.Arguments[2].(string)

		logs.Log("/source.shader/set/input", layer, index, value)
		dmn.SetShaderInput(layer, index, value)
	})
	server.Handle("/source.shader/set/uniform1f", func(msg *osc.Message) {
		layer := msg.Arguments[0].(string)
		name := msg.Arguments[1].(string)
		var value float32
		switch v := msg.Arguments[2].(type) {
		case int:
			value = float32(v)
		case int16:
			value = float32(v)
		case int32:
			value = float32(v)
		case int64:
			value = float32(v)
		case uint:
			value = float32(v)
		case uint16:
			value = float32(v)
		case uint32:
			value = float32(v)
		case uint64:
			value = float32(v)
		case float64:
			value = float32(v)
		case float32:
			value = v
		default:
			logs.Error("Expected to receive float32 uniform, but it was %T\n", value)
			return
		}

		logs.Log("/source.shader/set/uniform1f", layer, name, value)
		dmn.SetUniform(layer, name, value)
	})
	server.Handle("/program/create", func(msg *osc.Message) {
		name := msg.Arguments[0].(string)
		vShaderPath := msg.Arguments[1].(string)
		fShaderPath := msg.Arguments[2].(string)

		vShader, err := ioutil.ReadFile(vShaderPath)
		if err != nil {
			return
		}
		fShader, err := ioutil.ReadFile(fShaderPath)
		if err != nil {
			return
		}

		logs.Log("/program/create", name, string(vShader), string(fShader))
		dmn.AddProgram(name, string(vShader), string(fShader))
	})
	server.ListenAndServe()
}
