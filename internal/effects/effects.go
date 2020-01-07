package effects

import (
	"errors"

	"github.com/dianelooney/gggv/internal/opengl"
)

var ErrWindowShouldClose error = errors.New("Window should close")

type Interface interface {
	SetScene(s *opengl.Scene)
	Perform() (err error)
}

type effect struct {
	Scene *opengl.Scene
}

func (e *effect) SetScene(s *opengl.Scene) {
	e.Scene = s
}

type MaybeCloseWindow struct {
	effect
}

func (x MaybeCloseWindow) Perform() (err error) {
	if x.Scene.Window.ShouldClose() {
		return ErrWindowShouldClose
	}

	return
}

type CreateSourceFFT struct {
	effect
	Name string `osc:"0"`
}

func (x CreateSourceFFT) Perform() (err error) {
	x.Scene.AddSourceFFT(x.Name)
	return
}

type CreateSourceFFVideo struct {
	effect
	Name string
	Path string
}

func (x CreateSourceFFVideo) Perform() (err error) {
	x.Scene.AddSourceFFVideo(x.Name, x.Path)
	return
}

type CreateSourceShader struct {
	effect
	Name string `osc:"0"`
}

func (x CreateSourceShader) Perform() (err error) {
	x.Scene.AddSourceShader(x.Name)
	return
}

type CreateProgram struct {
	effect
	Name string
	Vert string
	Geom string
	Frag string
}

func (x CreateProgram) Perform() (err error) {
	x.Scene.LoadProgram(x.Name, x.Vert, x.Geom, x.Frag)
	return
}

type SetFFTScale struct {
	effect
	Name  string
	Scale float64
}

func (x SetFFTScale) Perform() (err error) {
	x.Scene.SetFFTScale(x.Name, float32(x.Scale))
	return
}

type SetFFVideoTimescale struct {
	effect
	Name  string
	Scale float64
}

func (x SetFFVideoTimescale) Perform() (err error) {
	x.Scene.SetFFVideoTimescale(x.Name, x.Scale)
	return
}

type SetShaderInput struct {
	effect
	Name   string
	Index  int
	Source string
}

func (x SetShaderInput) Perform() (err error) {
	x.Scene.SetShaderInput(x.Name, int32(x.Index), x.Source)
	return
}

type SetShaderProgram struct {
	effect
	Name    string
	Program string
}

func (x SetShaderProgram) Perform() (err error) {
	x.Scene.SetShaderProgram(x.Name, x.Program)
	return
}

type SetShaderUniformValue struct {
	effect
	Name  string
	Value interface{}
}

type SetShaderUniformClock struct {
	effect
	Name string
}

type SetWrapS struct {
	effect
	Name  string
	Value string
}

func (x SetWrapS) Perform() (err error) {
	x.Scene.SetSourceWrapS(x.Name, x.Value)
	return
}

type SetWrapT struct {
	effect
	Name  string
	Value string
}

func (x SetWrapT) Perform() (err error) {
	x.Scene.SetSourceWrapT(x.Name, x.Value)
	return
}

type SetMinFilter struct {
	effect
	Name  string
	Value string
}

func (x SetMinFilter) Perform() (err error) {
	x.Scene.SetSourceMinFilter(x.Name, x.Value)
	return
}

type SetMagFilter struct {
	effect
	Name  string
	Value string
}

func (x SetMagFilter) Perform() (err error) {
	x.Scene.SetSourceMagFilter(x.Name, x.Value)
	return
}
