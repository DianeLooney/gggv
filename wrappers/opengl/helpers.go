package opengl

import (
	"github.com/dianelooney/gggv/internal/carbon"
	"github.com/go-gl/gl/all-core/gl"
)

func init() {
	carbon.Ptr = gl.Ptr
	carbon.PtrOffset = gl.PtrOffset
	carbon.Str = gl.Str
	carbon.GoStr = gl.GoStr
	carbon.Strs = gl.Strs
}
