package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
	go func() {
		client := osc.NewClient("localhost", 4200)
		for {
			time.Sleep(time.Second)
			fmt.Println(client.Send(osc.NewMessage("/source.shader/set/source", "default", int32(0), "default0")))
			fmt.Println("source changed 0")
			time.Sleep(time.Second)
			fmt.Println(client.Send(osc.NewMessage("/source.shader/set/source", "default", int32(0), "default1")))
			fmt.Println("source changed 1")
			time.Sleep(time.Second)
			fmt.Println(client.Send(osc.NewMessage("/source.shader/set/source", "default", int32(0), "default2")))
			fmt.Println("source changed 2")
		}
	}()
	go func() {
		client := osc.NewClient("localhost", 4200)
		for {
			time.Sleep(time.Second * 2 / 3)
			fmt.Println(client.Send(osc.NewMessage("/source.shader/set/uniform1f", "default", "randR", -1+2*rand.Float32())))
			fmt.Println(client.Send(osc.NewMessage("/source.shader/set/uniform1f", "default", "randG", -1+2*rand.Float32())))
			fmt.Println(client.Send(osc.NewMessage("/source.shader/set/uniform1f", "default", "randB", -1+2*rand.Float32())))
			fmt.Println("uniform changed")
		}
	}()
	<-make(chan bool)
}
