package main

import (
	"flag"
	"fmt"
	_ "image/png"
	"io/ioutil"
	"runtime"
	"time"

	"github.com/dianelooney/gggv/internal/logs"
	"github.com/dianelooney/gggv/pkg/daemon"
	_ "github.com/dianelooney/gggv/wrappers/opengl" // necessary to fill carbon stubs
	"github.com/radovskyb/watcher"

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

	reloadProgram := func() {
		vShader, err := ioutil.ReadFile(*vShaderPath)
		if err != nil {
			return
		}
		gShader, err := ioutil.ReadFile(*gShaderPath)
		if err != nil {
			return
		}
		fShader, err := ioutil.ReadFile(*fShaderPath)
		if err != nil {
			return
		}
		dmn.Scene.LoadProgram("default", string(vShader), string(gShader), string(fShader))
	}
	reloadProgram()
	w := watcher.New()
	if err := w.Add(*vShaderPath); err != nil {
		logs.Log("Unable to watch shader", "default", vShaderPath, err)
		return
	}
	if err := w.Add(*gShaderPath); err != nil {
		logs.Log("Unable to watch shader", "default", gShaderPath, err)
		return
	}
	if err := w.Add(*fShaderPath); err != nil {
		logs.Log("Unable to watch shader", "default", fShaderPath, err)
		return
	}
	go func() {
		for {
			<-w.Event
			logs.Log("Reloading shader")
			reloadProgram()
		}
	}()
	go w.Start(time.Second / 3)

	dmn.Scene.BindBuffers()
	dmn.Scene.AddWindow()
	dmn.Scene.AddSourceShader("default")
	dmn.Scene.SetShaderInput("window", 0, "default")
	var sources = flag.Args()
	for i, s := range sources {
		fmt.Println("Adding source", i, s)
		dmn.Scene.AddSourceFFVideo(s, s)
		dmn.Scene.SetShaderInput("default", int32(i), s)
	}

	dmn.DrawLoop()
}

var pprof = flag.Bool("pprof", false, "Enable pprof on :6060 or not.")

var vShaderPath = flag.String("vert", "shaders/vert/default.glsl", "Path to the vertex shader to use")
var fShaderPath = flag.String("frag", "shaders/frag/default.glsl", "Path to the fragment shader to use")
var gShaderPath = flag.String("geom", "shaders/geom/default.glsl", "Path to the geom shader to use")

func enablePprof() {
	if *pprof {
		go func() {
			http.ListenAndServe("localhost:6060", nil)
		}()
	}
}
