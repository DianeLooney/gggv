package carbon_test

import (
	"fmt"
	"testing"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/dianelooney/gggv/internal/carbon"

	"github.com/dianelooney/gggv/wrappers/mock"
)

func TestUniform(t *testing.T) {
	mock.Init()

	f32Fixtures := []interface{}{
		float32(0),
		float64(0),
		int(0),
	}
	for _, fixture := range f32Fixtures {
		t.Run(fmt.Sprintf("It maps %T to Uniform1f", fixture), func(t *testing.T) {
			mock.Init()
			called := false
			carbon.Uniform1f = func(int32, float32) { called = true }
			carbon.Uniform(0, "", fixture)
			if !called {
				t.Fail()
			}
		})
	}

	t.Run("It map mgl32.Mat4 to UniformMatrix4fv", func(t *testing.T) {
		mock.Init()
		called := false
		carbon.UniformMatrix4fv = func(int32, int32, bool, *float32) { called = true }
		carbon.Uniform(0, "", mgl32.Mat4{})
		if !called {
			t.Error("UniformMatrix4fv was not called")
		}
	})
}
