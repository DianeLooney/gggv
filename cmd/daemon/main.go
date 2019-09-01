package main

import (
	"flag"
	_ "image/png"
	"io/ioutil"
	"runtime"

	"github.com/fsnotify/fsnotify"

	"github.com/dianelooney/gggv/internal/logs"

	"github.com/hypebeast/go-osc/osc"

	"github.com/dianelooney/gggv/pkg/daemon"
	_ "github.com/dianelooney/gggv/wrappers/opengl" // necessary to fill carbon stubs
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

	dmn.Scene.BindBuffers()
	dmn.AddSourceFFVideo("default0", "sample0.mp4")
	dmn.AddSourceFFVideo("default1", "sample1.mp4")
	dmn.AddSourceFFVideo("default2", "sample2.mp4")
	dmn.AddSourceShader("default")
	dmn.SetShaderInput("default", 0, "default0")
	dmn.SetShaderInput("default", 1, "default1")
	dmn.SetShaderInput("default", 2, "default2")
	dmn.Scene.AddWindow()
	dmn.SetShaderInput("window", 0, "default")
	dmn.SetUniform("default", "ampl", float32(0.0))

	go netSetup()
	dmn.DrawLoop()
}

var netAddr = flag.String("net", ":4200", "Network address to listen at.")

func netSetup() {
	server := &osc.Server{Addr: *netAddr}
	server.Handle("/source/set/wrap.s", func(msg *osc.Message) {
		name := msg.Arguments[0].(string)
		opt := msg.Arguments[1].(string)
		dmn.SetSourceWrapS(name, opt)
	})
	server.Handle("/source/set/wrap.t", func(msg *osc.Message) {
		name := msg.Arguments[0].(string)
		opt := msg.Arguments[1].(string)
		dmn.SetSourceWrapT(name, opt)
	})
	server.Handle("/source/set/minfilter", func(msg *osc.Message) {
		name := msg.Arguments[0].(string)
		opt := msg.Arguments[1].(string)
		dmn.SetSourceMinFilter(name, opt)
	})
	server.Handle("/source/set/magfilter", func(msg *osc.Message) {
		name := msg.Arguments[0].(string)
		opt := msg.Arguments[1].(string)
		dmn.SetSourceMagFilter(name, opt)
	})
	//TODO: better error handling on all of the type assertions.
	server.Handle("/source.ffvideo/create", func(msg *osc.Message) {
		sourceName := msg.Arguments[0].(string)
		path := msg.Arguments[1].(string)

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
	watchers := make(map[string]chan bool)
	loadProgram := func(name, vShaderPath, fShaderPath string) {
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
	server.Handle("/program/create", func(msg *osc.Message) {
		name := msg.Arguments[0].(string)
		vShaderPath := msg.Arguments[1].(string)
		fShaderPath := msg.Arguments[2].(string)
		loadProgram(name, vShaderPath, fShaderPath)

		if ch, ok := watchers[name]; ok {
			ch <- true
			delete(watchers, name)
		}
		logs.Log("/program/create", name, vShaderPath, fShaderPath)
	})
	server.Handle("/program/watch", func(msg *osc.Message) {
		name := msg.Arguments[0].(string)
		vShaderPath := msg.Arguments[1].(string)
		fShaderPath := msg.Arguments[2].(string)

		loadProgram(name, vShaderPath, fShaderPath)

		if ch, ok := watchers[name]; ok {
			ch <- true
			delete(watchers, name)
		}

		done := make(chan bool)
		watchers[name] = done
		w, err := fsnotify.NewWatcher()
		if err != nil {
			return
		}
		w.Add(vShaderPath)
		w.Add(fShaderPath)

		go func() {
			defer logs.Log("watcher killed", name, vShaderPath, fShaderPath)
			for {
				select {
				case e := <-w.Events:
					logs.Log("reloading", name, vShaderPath, fShaderPath, e)
					loadProgram(name, vShaderPath, fShaderPath)
				case <-done:
					w.Close()
					return
				}
			}
		}()
		logs.Log("/program/watch", name, vShaderPath, fShaderPath)
	})
	server.ListenAndServe()
}
