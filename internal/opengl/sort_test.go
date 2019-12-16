package opengl_test

import (
	"testing"

	. "github.com/dianelooney/gggv/internal/opengl"
)

type testSource struct {
	children []SourceName
}

func (s testSource) Name() SourceName {
	return ""
}
func (s testSource) Children() []SourceName {
	return s.children
}
func (s testSource) Render(scene *Scene) {
	return
}
func (s testSource) SkipRender(scene *Scene) {
	return
}
func (s testSource) Dimensions() (width, height int32) {
	return 0, 0
}
func (s testSource) Texture() uint32 {
	return 0
}

func TestOrder(t *testing.T) {
	t.Run("Doesnt fail on a simple case", func(t *testing.T) {
		m := make(map[SourceName]Source)
		m[SourceName("a")] = testSource{[]SourceName{"b", "c"}}
		m[SourceName("b")] = testSource{}
		m[SourceName("c")] = testSource{}
		_, err := Order(SourceName("a"), m)
		if err != nil {
			t.Errorf("err was not nil: %v", err)
		}
	})

	t.Run("Doesnt duplicate render calls", func(t *testing.T) {
		m := make(map[SourceName]Source)
		m[SourceName("a")] = testSource{[]SourceName{"b", "c"}}
		m[SourceName("b")] = testSource{[]SourceName{"c"}}
		m[SourceName("c")] = testSource{}
		ord, err := Order(SourceName("a"), m)
		if err != nil {
			t.Errorf("err was not nil: %v", err)
		}
		fixture := "cba"
		var actual string
		for _, s := range ord {
			actual += string(s)
		}
		if fixture != actual {
			t.Errorf("render order mismatch\nfixture: %v\ntest: %v", fixture, actual)
		}
	})

	t.Run("Fails when there is a loop", func(t *testing.T) {
		m := make(map[SourceName]Source)
		m[SourceName("a")] = testSource{[]SourceName{"b", "c"}}
		m[SourceName("b")] = testSource{}
		m[SourceName("c")] = testSource{[]SourceName{"a"}}
		_, err := Order(SourceName("a"), m)
		if err == nil {
			t.Errorf("err was nil: %v", err)
		}
	})

	t.Run("Fails when there is a loop", func(t *testing.T) {
		m := make(map[SourceName]Source)
		m[SourceName("a")] = testSource{[]SourceName{"b"}}
		m[SourceName("b")] = testSource{[]SourceName{"c"}}
		m[SourceName("c")] = testSource{[]SourceName{"a"}}
		_, err := Order(SourceName("a"), m)
		if err == nil {
			t.Errorf("err was nil: %v", err)
		}
	})

	t.Run("Fails when sources are missing", func(t *testing.T) {
		m := make(map[SourceName]Source)
		m[SourceName("a")] = testSource{[]SourceName{"b", "c"}}
		m[SourceName("b")] = testSource{}
		_, err := Order(SourceName("a"), m)
		if err == nil {
			t.Errorf("err was nil: %v", err)
		}
	})
}
