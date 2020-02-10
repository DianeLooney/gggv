package shader

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/go-gl/glfw/v3.1/glfw"

	"github.com/dianelooney/gggv/internal/carbon"
	_ "github.com/dianelooney/gggv/wrappers/opengl"
)

func TestShader(t *testing.T) {
	if err := glfw.Init(); err != nil {
		t.Errorf("Unable to initialize glfw: %v", err)
		return
	}
	glfw.WindowHint(glfw.Visible, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 5)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	w, err := glfw.CreateWindow(1920, 1080, "gggv-test", nil, nil)
	if err != nil {
		t.Errorf("Unable to create glfw window: %v", err)
		return
	}
	w.MakeContextCurrent()
	if err := carbon.Init(); err != nil {
		t.Errorf("Unable to initialize carbon: %v", err)
		return
	}

	for name, kind := range map[string]uint32{
		"frag": carbon.FRAGMENT_SHADER,
		"geom": carbon.GEOMETRY_SHADER,
		"vert": carbon.VERTEX_SHADER,
	} {
		t.Run(name, func(t *testing.T) {
			f, err := os.OpenFile(name, os.O_RDONLY, 0744)
			if err != nil {
				t.Errorf("Unable to open %s shader directory: %v", name, err)
				return
			}
			if f == nil {
				t.Errorf("%s directory was nil", name)
				return
			}

			names, err := f.Readdirnames(0)
			if err != nil {
				t.Errorf("Unable to enumerate %s shader directory: %v", name, err)
				return
			}

			stubData, err := ioutil.ReadFile(name + ".stub.glsl")
			if err != nil {
				t.Errorf("Unable to read %s stub: %v", name, err)
				return
			}

			for _, n := range names {
				t.Run(n, func(t *testing.T) {
					fData, err := ioutil.ReadFile(name + "/" + n)
					if err != nil {
						t.Error(err)
						return
					}
					sh, err := carbon.WrappedCompileShader(string(stubData)+string(fData)+"\x00", kind)
					if err != nil {
						t.Error(err)
						return
					}
					carbon.DeleteShader(sh)
				})
			}
		})
	}
}
