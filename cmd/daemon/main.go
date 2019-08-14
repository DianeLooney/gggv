package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "image/png"
	"io"
	"net"
	"runtime"

	"github.com/dianelooney/gvd/internal/opengl"

	"github.com/dianelooney/gvd/pkg/gvdl"

	"github.com/dianelooney/gvd/pkg/daemon"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var dmn *daemon.D

func main() {
	flag.Parse()
	netSetup()

	dmn = daemon.New()

	if err := dmn.Scene.LoadProgram("default", "shaders/vert/default.glsl", "shaders/frag/default.glsl"); err != nil {
		panic(err)
	}

	if err := dmn.Scene.LoadProgram("final", "shaders/vert/final.glsl", "shaders/frag/final.glsl"); err != nil {
		panic(err)
	}

	dmn.Scene.SetLayer("default", 0.0, "default", [opengl.LAYER_TEXTURE_COUNT]string{"default0", "default1", "default2"})

	dmn.Scene.BindBuffers()
	dmn.AddSource("default0", "sample0.mp4")
	dmn.AddSource("default1", "sample1.mp4")
	dmn.AddSource("default2", "sample2.mp4")

	dmn.DrawLoop()
}

var netAddr = flag.String("net", ":4200", "Network address to listen at.")

func netSetup() {
	srv, err := net.Listen("tcp", *netAddr)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			conn, err := srv.Accept()
			if err != nil {
				panic(err)
			}

			go handleConnection(conn)
		}
	}()
}

func handleConnection(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("Error reading from net:", err)
			continue
		}
		err = gvdl.Exec(line, dmn)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
