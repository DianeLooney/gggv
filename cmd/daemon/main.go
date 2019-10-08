package main

import (
	"flag"
	_ "image/png"
	"io/ioutil"
	"runtime"
	"time"

	"github.com/radovskyb/watcher"

	"github.com/dianelooney/gggv/internal/logs"
	"github.com/dianelooney/gggv/internal/net"

	"github.com/hypebeast/go-osc/osc"

	"github.com/dianelooney/gggv/pkg/daemon"
	_ "github.com/dianelooney/gggv/wrappers/opengl" // necessary to fill carbon stubs

	"net/http"
	_ "net/http/pprof" // makes things go fast
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var dmn *daemon.D

func main() {
	flag.Parse()
	enablePprof()

	dmn = daemon.New()

	{
		name := "default"
		vShaderPath := "shaders/vert/default.glsl"
		gShaderPath := "shaders/geom/default.glsl"
		fShaderPath := "shaders/frag/default.glsl"
		vShader, err := ioutil.ReadFile(vShaderPath)
		if err != nil {
			return
		}
		fShader, err := ioutil.ReadFile(fShaderPath)
		if err != nil {
			return
		}
		gShader, err := ioutil.ReadFile(gShaderPath)
		if err != nil {
			return
		}
		dmn.AddProgram(name, string(vShader), string(gShader), string(fShader))
	}

	dmn.Scene.BindBuffers()
	dmn.Scene.AddWindow()

	go netSetup()
	dmn.DrawLoop()
}

var pprof = flag.Bool("pprof", false, "Enable pprof on :6060 or not.")

func enablePprof() {
	if *pprof {
		go func() {
			http.ListenAndServe("localhost:6060", nil)
		}()
	}
}

var addr = flag.String("net", ":4200", "Network address to listen at.")

func netSetup() {
	server := &osc.Server{Addr: *addr}

	net.HandleSS(server, "/source/set/wrap.s", "name", "opt", dmn.SetSourceWrapS)
	net.HandleSS(server, "/source/set/wrap.t", "name", "opt", dmn.SetSourceWrapT)
	net.HandleSS(server, "/source/set/minfilter", "name", "opt", dmn.SetSourceMinFilter)
	net.HandleSS(server, "/source/set/magfilter", "name", "opt", dmn.SetSourceMagFilter)
	net.HandleSS(server, "/source.ffvideo/create", "name", "opt", dmn.AddSourceFFVideo)
	net.HandleSF(server, "/source.ffvideo/set/timescale", "name", "timescale", dmn.SetFFVideoTimescale)
	net.HandleS(server, "/source.shader/create", "name", dmn.AddSourceShader)
	net.HandleSIS(server, "/source.shader/set/input", "name", "index", "value", dmn.SetShaderInput)
	net.HandleSS(server, "/source.shader/set/program", "shader", "program", dmn.SetShaderProgram)
	net.HandleSSFFF(server, "/source.shader/set/uniform3f", "shader", "uniform", "vec[0]", "vec[1]", "vec[2]", dmn.SetUniform3f)
	net.HandleSS(server, "/source.shader/set/uniform.clock", "shader", "uniform", dmn.SetUniformClock)
	net.HandleSS(server, "/source.shader/set/uniform.timestamp", "shader", "uniform", dmn.SetUniformTimestamp)

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
	loadProgram := func(name, vShaderPath, gShaderPath, fShaderPath string) {
		vShader, err := ioutil.ReadFile(vShaderPath)
		if err != nil {
			return
		}
		gShader, err := ioutil.ReadFile(gShaderPath)
		if err != nil {
			return
		}
		fShader, err := ioutil.ReadFile(fShaderPath)
		if err != nil {
			return
		}
		dmn.AddProgram(name, string(vShader), string(gShader), string(fShader))
	}
	server.Handle("/program/create", func(msg *osc.Message) {
		name := msg.Arguments[0].(string)
		vShaderPath := msg.Arguments[1].(string)
		gShaderPath := msg.Arguments[2].(string)
		fShaderPath := msg.Arguments[3].(string)
		loadProgram(name, vShaderPath, gShaderPath, fShaderPath)

		if ch, ok := watchers[name]; ok {
			ch <- true
			delete(watchers, name)
		}
		logs.Log("/program/create", name, vShaderPath, fShaderPath)
	})
	server.Handle("/program/watch", func(msg *osc.Message) {
		name := msg.Arguments[0].(string)
		vShaderPath := msg.Arguments[1].(string)
		gShaderPath := msg.Arguments[2].(string)
		fShaderPath := msg.Arguments[3].(string)

		loadProgram(name, vShaderPath, gShaderPath, fShaderPath)

		if ch, ok := watchers[name]; ok {
			ch <- true
			delete(watchers, name)
		}

		done := make(chan bool)
		watchers[name] = done

		w := watcher.New()
		if err := w.Add(vShaderPath); err != nil {
			logs.Log("Unable to watch shader", name, vShaderPath, err)
			return
		}
		if err := w.Add(gShaderPath); err != nil {
			logs.Log("Unable to watch shader", name, gShaderPath, err)
			return
		}
		if err := w.Add(fShaderPath); err != nil {
			logs.Log("Unable to watch shader", name, fShaderPath, err)
			return
		}
		go func() {
			for {
				select {
				case e := <-w.Event:
					logs.Log("Reloading shader", name, vShaderPath, gShaderPath, fShaderPath, e)
					loadProgram(name, vShaderPath, gShaderPath, fShaderPath)
				case <-done:
					w.Close()
					return
				}
			}
		}()
		go w.Start(time.Second / 3)

		logs.Log("/program/watch", name, vShaderPath, fShaderPath)
	})
	server.ListenAndServe()
}
