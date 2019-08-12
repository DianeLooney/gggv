package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "image/png"
	"io"
	"log"
	"os"
	"runtime"
	"syscall"

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
	pipeSetup()

	dmn = daemon.New()

	if err := dmn.Scene.LoadProgram("default", "shaders/vert/default.glsl", "shaders/frag/default.glsl"); err != nil {
		panic(err)
	}

	dmn.Scene.SetLayer("default", 0.0, "default")
	//dmn.Scene.SetLayer("sub", -1.0, "swap")

	dmn.Scene.BindBuffers()
	dmn.Scene.TextureInit("default")
	dmn.AddSource("default", "sample.mp4")
	//dmn.AddSource("swap", "sample2.mp4")

	dmn.DrawLoop()
}

var pipeFile = flag.String("pipe", "tmp/daemon.pipe", "Path to a named pipe to use for communication")

func pipeSetup() {
	fmt.Println("remove")
	os.Remove(*pipeFile)
	fmt.Println("mkfifo")
	err := syscall.Mkfifo(*pipeFile, 0666)
	if err != nil {
		log.Fatal("Unable to create named pipe:", err)
	}

	go func() {
		for {
			file, err := os.OpenFile(*pipeFile, os.O_CREATE, 0600)
			if err != nil {
				log.Fatal("Unable to open named pipe:", err)
			}
			handlePipe(bufio.NewReader(file), bufio.NewWriter(file))
		}
	}()
}

func handlePipe(r *bufio.Reader, w *bufio.Writer) {
	fmt.Println("Client connected")
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			fmt.Println("Client disconnected")
		}
		if err != nil {
			fmt.Println("Error reading from pipe:", err)
			continue
		}
		err = gvdl.Exec(line, dmn)
		if err != nil {
			fmt.Fprintln(w, "Error:", err)
		}
	}
}
