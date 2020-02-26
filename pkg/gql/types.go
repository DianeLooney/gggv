package gql

import (
	"hash/fnv"

	"github.com/graph-gophers/graphql-go"

	"github.com/dianelooney/gggv/internal/opengl"
	"github.com/dianelooney/gggv/pkg/daemon"
)

func hash(text string) int {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(text))
	return int(algorithm.Sum32())
}

type Schema struct {
	d *daemon.D
}

type Source struct {
	s opengl.Source
}

func (s Source) ID() graphql.ID { return graphql.ID(s.s.Name()) }
func (s Source) SourceIDs() []graphql.ID {
	out := make([]graphql.ID, len(s.s.Children()))
	for i, c := range s.s.Children() {
		out[i] = graphql.ID(c)
	}
	return out
}
func (s Source) Uniforms() *[]Uniform {
	sh, ok := s.s.(*opengl.ShaderSource)
	if !ok {
		return nil
	}

	out := make([]Uniform, 0)
	for _, u := range sh.Uniforms {
		out = append(out, Uniform{u})
	}
	return &out
}
func (s Source) ProgramID() *graphql.ID {
	sh, ok := s.s.(*opengl.ShaderSource)
	if !ok {
		return nil
	}
	v := graphql.ID(sh.P)
	return &v
}

type Uniform struct {
	v opengl.BindUniformer
}

func (u Uniform) Name() string {
	switch x := u.v.(type) {
	case opengl.ValueUniform:
		return x.Name
	case opengl.ClockUniform:
		return x.Name
	}
	panic("NYI")
}
func (u Uniform) Value1f() *float64 {
	switch x := u.v.(type) {
	case opengl.ValueUniform:
		v, ok := x.Value.(float32)
		if !ok {
			return nil
		}
		v64 := float64(v)
		return &v64
	}
	panic("NYI")
}
