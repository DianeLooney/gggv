package net

import (
	"flag"

	"github.com/dianelooney/gggv/internal/errors"
	"github.com/dianelooney/gggv/internal/logs"

	"github.com/hypebeast/go-osc/osc"
)

var verbose = flag.Bool("v", false, "Verbose output, logs osc endpoints")

func Handle(d *osc.StandardDispatcher, path string, f Handler) {
	d.AddMsgHandler(path, func(msg *osc.Message) {
		defer func() {
			if err := recover(); err != nil {
				logs.Log("recovered from panic", path, err)
			}
		}()

		args := &arguments{
			endpoint: path,
			offset:   0,
			msg:      msg,
		}
		if *verbose {
			logs.Log(path, msg.Arguments)
		}
		f(args)
	})
}

type Handler func(args Shifter)

type Shifter interface {
	Length() int
	Shift() interface{}
}

type arguments struct {
	endpoint string
	offset   int
	msg      *osc.Message
}

func (p *arguments) Shift() interface{} {
	defer func() { p.offset++ }()
	length := p.msg.CountArguments()
	if p.offset >= length {
		panic(errors.NetOutOfArgs(p.offset))
	}
	return p.msg.Arguments[p.offset]
}

func (p *arguments) Length() int {
	return p.msg.CountArguments()
}
