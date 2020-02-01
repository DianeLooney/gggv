package carbon

import (
	"strings"

	"github.com/dianelooney/gggv/internal/errors"
	"github.com/go-gl/mathgl/mgl32"
)

func WrappedCompileShader(source string, shaderType uint32) (uint32, error) {
	shader := CreateShader(shaderType)

	csources, free := Strs(source)
	ShaderSource(shader, 1, csources, nil)
	free()
	CompileShader(shader)

	var status int32
	GetShaderiv(shader, COMPILE_STATUS, &status)
	if status == FALSE {
		var logLength int32
		GetShaderiv(shader, INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		GetShaderInfoLog(shader, logLength, nil, Str(log))

		return 0, errors.GLShaderCompile(log, source)
	}

	return shader, nil
}

func Uniform(program uint32, name string, value interface{}) {
	u := GetUniformLocation(program, Str(name+"\x00"))
	switch v := value.(type) {
	case float32:
		Uniform1f(u, v)
	case float64:
		Uniform1f(u, float32(v))
	case int:
		Uniform1f(u, float32(v))
	case mgl32.Mat4:
		UniformMatrix4fv(u, 1, false, &v[0])
	case [2]float32:
		Uniform2f(u, v[0], v[1])
	case [3]float32:
		Uniform3f(u, v[0], v[1], v[2])
	case [4]float32:
		Uniform4f(u, v[0], v[1], v[2], v[3])
	}
}

func UniformTex(program uint32, name string, value int32) {
	u := GetUniformLocation(program, Str(name+"\x00"))
	Uniform1i(u, value)
}
