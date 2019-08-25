package main

import (
	"fmt"
	"time"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
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
}
