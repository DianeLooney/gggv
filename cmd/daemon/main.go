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
	dmn.Scene.AddWindow()

	go netSetup()

	dmn.Begin()
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
	net.RouteOSC(server, dmn)
	server.ListenAndServe()
}
