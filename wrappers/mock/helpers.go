package mock

import (
	"unsafe"

	"github.com/dianelooney/gggv/internal/carbon"
)

func initHelpers() {
	carbon.Ptr = func(data interface{}) unsafe.Pointer {
		return nil
	}
	carbon.PtrOffset = func(offset int) unsafe.Pointer {
		return nil
	}
	carbon.Str = func(str string) *uint8 {
		return nil
	}
	carbon.GoStr = func(cstr *uint8) string {
		return ""
	}
	carbon.Strs = func(strs ...string) (cstrs **uint8, free func()) {
		return nil, func() {}
	}
}
