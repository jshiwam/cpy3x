package py35

// #cgo pkg-config: python3
// #include "Python.h"
import "C"
import (
	"unsafe"

	"github.com/jshiwam/cpy3x/pycore"
)

// Part of Stable ABI
// PyBool_FromLong : https://docs.python.org/3.5/c-api/bool.html?highlight=pybool_type#c.PyBool_FromLong
func PyBool_FromLong(v int) *pycore.PyObject {
	return pycore.Togo(pycore.VoidToC(unsafe.Pointer(C.PyBool_FromLong(C.long(v)))))
}
