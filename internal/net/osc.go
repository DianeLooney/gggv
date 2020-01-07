package net

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/dianelooney/gggv/internal/effects"
	"github.com/dianelooney/gggv/pkg/daemon"
	"github.com/hypebeast/go-osc/osc"
)

func RouteOSC(s *osc.Server, d *daemon.D) {
	OSCBindRoute(s, d, "/source.shader/create", effects.CreateSourceShader{})
}

func OSCBindRoute(s *osc.Server, d *daemon.D, path string, val interface{}) {
	t := reflect.ValueOf(val).Type()
	s.Handle(path, func(msg *osc.Message) {
		effect := reflect.New(t)
		for i := 0; i < t.NumField(); i++ {
			sField := t.Field(i)
			tag, ok := sField.Tag.Lookup("osc")
			if !ok {
				continue
			}
			j, err := strconv.Atoi(strings.Split(tag, ",")[0])
			if err != nil {
				panic("Invalid struct tag " + tag)
			}
			v := msg.Arguments[j]
			vCasted := reflect.ValueOf(v).Convert(sField.Type)
			effect.Elem().Field(i).Set(vCasted)
		}
		iface := effect.Interface().(effects.Interface)
		iface.SetScene(d.Scene)
		d.Push(iface)
	})
	t.NumField()
}
