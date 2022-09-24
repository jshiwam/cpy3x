package pycore

/*
#include "Python.h"
*/
import "C"

import (
	"unsafe"
)

//Part Of Stable ABI
//PyImport_ImportModule : https://docs.python.org/3/c-api/import.html#c.PyImport_ImportModule
func PyImport_ImportModule(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Togo(C.PyImport_ImportModule(cname))
}
