package carbon

import "unsafe"

var Ptr func(data interface{}) unsafe.Pointer
var PtrOffset func(offset int) unsafe.Pointer
var Str func(str string) *uint8
var GoStr func(cstr *uint8) string
var Strs func(strs ...string) (cstrs **uint8, free func())
