package carbon

import "github.com/go-gl/mathgl/mgl32"

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
	case [3]float32:
		Uniform3f(u, v[0], v[1], v[2])
	}
}

func UniformTex(program uint32, name string, value int32) {
	u := GetUniformLocation(program, Str(name+"\x00"))
	Uniform1i(u, value)
}
