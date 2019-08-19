package main

import "github.com/hypebeast/go-osc/osc"

func main() {
	client := osc.NewClient("localhost", 4200)
	msg := osc.NewMessage("/programs/reload")
	client.Send(msg)
}
