package main

import (
	"flag"
	"fmt"
	_ "image/png"
	"runtime"

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

	dmn.Scene.BindBuffers()
	dmn.WatchProgramS("window", "shaders/vert/window.glsl", "shaders/geom/window.glsl", "shaders/frag/window.glsl")
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
	d := osc.NewStandardDispatcher()
	net.Handle(d, "/pause", dmn.Pause)
	net.Handle(d, "/play", dmn.Play)
	net.Handle(d, "/source/set/wrap.s", dmn.SetSourceWrapS)
	net.Handle(d, "/source/set/wrap.t", dmn.SetSourceWrapT)
	net.Handle(d, "/source/set/minfilter", dmn.SetSourceMinFilter)
	net.Handle(d, "/source/set/magfilter", dmn.SetSourceMagFilter)
	net.Handle(d, "/source.fft/create", dmn.AddSourceFFT)
	net.Handle(d, "/source.fft/scale", dmn.SetFFTScale)
	net.Handle(d, "/source.ffvideo/create", dmn.AddSourceFFVideo)
	net.Handle(d, "/source.ffvideo/set/timescale", dmn.SetFFVideoTimescale)
	net.Handle(d, "/source.shader/create", dmn.AddSourceShader)
	net.Handle(d, "/source.shader/set/input", dmn.SetShaderInput)
	net.Handle(d, "/source.shader/set/program", dmn.SetShaderProgram)
	net.Handle(d, "/source.shader/set/geometry", dmn.SetShaderGeometry)
	net.Handle(d, "/source.shader/set/drawcount", dmn.SetShaderDrawCount)
	net.Handle(d, "/source.shader/set/dimensions", dmn.SetShaderDimensions)
	net.Handle(d, "/source.shader/set/uniform1f", dmn.SetUniform)
	net.Handle(d, "/source.shader/set/uniform3f", dmn.SetUniform3f)
	net.Handle(d, "/source.shader/set/uniform.clock", dmn.SetUniformClock)
	net.Handle(d, "/source.shader/set/uniform.timestamp", dmn.SetUniformTimestamp)
	net.Handle(d, "/source.shader/set.global/uniform1f", dmn.SetUniform)
	net.Handle(d, "/source.shader/set.global/uniform3f", dmn.SetGlobalUniform3f)
	net.Handle(d, "/source.shader/set.global/uniform.clock", dmn.SetGlobalUniformClock)
	net.Handle(d, "/source.shader/set.global/uniform.timestamp", dmn.SetGlobalUniformTimestamp)
	net.Handle(d, "/program/create", dmn.CreateProgram)
	net.Handle(d, "/program/watch", dmn.WatchProgram)

	server := &osc.Server{Addr: *addr, Dispatcher: d}
	go func() {
		for {
			err := server.ListenAndServe()
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}
