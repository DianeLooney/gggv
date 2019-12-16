package net

import (
	"github.com/dianelooney/gggv/internal/errors"
	"github.com/dianelooney/gggv/internal/logs"

	"github.com/hypebeast/go-osc/osc"
)

type HanlderSS func(s1, s2 string)

func HandleSS(s *osc.Server, path string, n1 string, n2 string, f HanlderSS) {
	s.Handle(path, func(msg *osc.Message) {
		if count := msg.CountArguments(); count != 2 {
			logs.Error(errors.NetTooManyArgs(2, count))
			return
		}
		args := newArguments(path, msg)
		s1, err := args.NextString()
		if err != nil {
			logs.Error(path, n1, err)
			return
		}

		s2, err := args.NextString()
		if err != nil {
			logs.Error(path, n2, err)
			return
		}
		logs.Log(path, s1, s2)
		f(s1, s2)
	})
}

type HanlderSIS func(s1 string, i int32, s2 string)

func HandleSIS(s *osc.Server, path string, n1 string, n2 string, n3 string, f HanlderSIS) {
	s.Handle(path, func(msg *osc.Message) {
		logs.Log(path)
		if count := msg.CountArguments(); count != 3 {
			logs.Error(errors.NetTooManyArgs(3, count))
			return
		}
		args := newArguments(path, msg)
		s1, err := args.NextString()
		if err != nil {
			logs.Error(path, n1, err)
			return
		}

		i, err := args.NextInt32()
		if err != nil {
			logs.Error(path, n1, err)
			return
		}

		s2, err := args.NextString()
		if err != nil {
			logs.Error(path, n2, err)
			return
		}

		logs.Log(path, s1, i, s2)
		f(s1, i, s2)
	})
}

type HandlerSSFFF = func(s1, s2 string, v0, v1, v2 float32)

func HandleSSFFF(s *osc.Server, path string, n1, n2, n3, n4, n5 string, f HandlerSSFFF) {
	s.Handle(path, func(msg *osc.Message) {
		logs.Log(path)
		if count := msg.CountArguments(); count != 5 {
			logs.Error(errors.NetTooManyArgs(5, count))
			return
		}
		args := newArguments(path, msg)
		s1, err := args.NextString()
		if err != nil {
			logs.Error(path, n1, err)
			return
		}
		s2, err := args.NextString()
		if err != nil {
			logs.Error(path, n2, err)
			return
		}

		v0, err := args.NextFloat32()
		if err != nil {
			logs.Error(path, n1, err)
			return
		}

		v1, err := args.NextFloat32()
		if err != nil {
			logs.Error(path, n1, err)
			return
		}

		v2, err := args.NextFloat32()
		if err != nil {
			logs.Error(path, n1, err)
			return
		}

		logs.Log(path, s1, s2, v0, v1, v2)
		f(s1, s2, v0, v1, v2)
	})
}

type HandlerSF = func(s string, v float32)

func HandleSF(s *osc.Server, path string, n1, n2 string, f HandlerSF) {
	s.Handle(path, func(msg *osc.Message) {
		logs.Log(path)
		if count := msg.CountArguments(); count != 2 {
			logs.Error(errors.NetTooManyArgs(2, count))
			return
		}
		args := newArguments(path, msg)
		s, err := args.NextString()
		if err != nil {
			logs.Error(path, n1, err)
			return
		}

		v, err := args.NextFloat32()
		if err != nil {
			logs.Error(path, n1, err)
			return
		}

		logs.Log(path, s, v)
		f(s, v)
	})
}

type HanlderS func(s string)

func HandleS(s *osc.Server, path string, n string, f HanlderS) {
	s.Handle(path, func(msg *osc.Message) {
		if count := msg.CountArguments(); count != 1 {
			logs.Error(errors.NetTooManyArgs(1, count))
			return
		}
		args := newArguments(path, msg)
		s, err := args.NextString()
		if err != nil {
			logs.Error(path, n, err)
			return
		}
		logs.Log(path, s)
		f(s)
	})
}

func newArguments(endpoint string, msg *osc.Message) *Arguments {
	return &Arguments{
		endpoint: endpoint,
		offset:   0,
		msg:      msg,
	}
}

type Arguments struct {
	endpoint string
	offset   int
	msg      *osc.Message
}

func (p *Arguments) NextString() (string, error) {
	defer func() { p.offset++ }()

	length := p.msg.CountArguments()
	if p.offset >= length {
		return "", errors.NetOutOfArgs(p.offset, "string")
	}

	v := p.msg.Arguments[p.offset]
	s, ok := v.(string)
	if !ok {
		return "", errors.NetWrongArgType(p.offset, "string", v)
	}

	return s, nil
}

func (p *Arguments) NextInt32() (int32, error) {
	defer func() { p.offset++ }()

	length := p.msg.CountArguments()
	if p.offset >= length {
		return 0, errors.NetOutOfArgs(p.offset, "int32")
	}

	v := p.msg.Arguments[p.offset]
	f, ok := v.(int32)
	if !ok {
		return 0, errors.NetWrongArgType(p.offset, "int32", v)
	}

	return f, nil
}

func (p *Arguments) NextFloat32() (float32, error) {
	defer func() { p.offset++ }()

	length := p.msg.CountArguments()
	if p.offset >= length {
		return 0, errors.NetOutOfArgs(p.offset, "float32")
	}

	v := p.msg.Arguments[p.offset]

	if f, ok := v.(float32); ok {
		return f, nil
	}

	if i, ok := v.(int32); ok {
		return float32(i), nil
	}

	return 0, errors.NetWrongArgType(p.offset, "float32", v)
}
