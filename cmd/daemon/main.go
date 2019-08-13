package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "image/png"
	"io"
	"net"
	"runtime"

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

	dmn.Scene.SetLayer("default", 0.0, "default", "default")
	//dmn.Scene.SetLayer("sub", -1.0, "swap")

	dmn.Scene.BindBuffers()
	dmn.AddSource("default", "sample.mp4")

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
