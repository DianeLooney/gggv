package main

import (
	"flag"
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
	server := &osc.Server{Addr: *addr}

	net.Handle(server, "/pause", dmn.Pause)
	net.Handle(server, "/play", dmn.Play)
	net.Handle(server, "/source/set/wrap.s", dmn.SetSourceWrapS)
	net.Handle(server, "/source/set/wrap.t", dmn.SetSourceWrapT)
	net.Handle(server, "/source/set/minfilter", dmn.SetSourceMinFilter)
	net.Handle(server, "/source/set/magfilter", dmn.SetSourceMagFilter)
	net.Handle(server, "/source.fft/create", dmn.AddSourceFFT)
	net.Handle(server, "/source.fft/scale", dmn.SetFFTScale)
	net.Handle(server, "/source.ffvideo/create", dmn.AddSourceFFVideo)
	net.Handle(server, "/source.ffvideo/set/timescale", dmn.SetFFVideoTimescale)
	net.Handle(server, "/source.shader/create", dmn.AddSourceShader)
	net.Handle(server, "/source.shader/set/input", dmn.SetShaderInput)
	net.Handle(server, "/source.shader/set/program", dmn.SetShaderProgram)
	net.Handle(server, "/source.shader/set/uniform1f", dmn.SetUniform)
	net.Handle(server, "/source.shader/set/uniform3f", dmn.SetUniform3f)
	net.Handle(server, "/source.shader/set/uniform.clock", dmn.SetUniformClock)
	net.Handle(server, "/source.shader/set/uniform.timestamp", dmn.SetUniformTimestamp)
	net.Handle(server, "/source.shader/set.global/uniform1f", dmn.SetUniform)
	net.Handle(server, "/source.shader/set.global/uniform3f", dmn.SetGlobalUniform3f)
	net.Handle(server, "/source.shader/set.global/uniform.clock", dmn.SetGlobalUniformClock)
	net.Handle(server, "/source.shader/set.global/uniform.timestamp", dmn.SetGlobalUniformTimestamp)
	net.Handle(server, "/program/create", dmn.CreateProgram)
	net.Handle(server, "/program/watch", dmn.WatchProgram)
	server.ListenAndServe()
}
