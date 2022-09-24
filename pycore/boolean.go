package pycore

// #cgo pkg-config: python3
// #include "Python.h"
// #include "macros.h"
import "C"
import "unsafe"

// Part Of Stable ABI
// PyBool_Type : https://docs.python.org/3/c-api/stable.html
var Bool = Togo((*C.PyObject)(unsafe.Pointer(&C.PyBool_Type)))

// Not part of Stable ABI, but available across versions 3.5-3.10
// Py_False : https://docs.python.org/3.5/c-api/bool.html?highlight=pybool_type#c.Py_False
var (
	Py_False = Togo(C.Py_False)
	Py_True  = Togo(C.Py_True)
)

// Part of Stable ABI
// PyBool_FromLong : https://docs.python.org/3.5/c-api/bool.html?highlight=pybool_type#c.PyBool_FromLong
func PyBool_FromLong(v int) *PyObject {
	return Togo(C.PyBool_FromLong(C.long(v)))
}

// Not part of Stable ABI, but available across versions 3.5-3.10
// PyBool_Check : https://docs.python.org/3.10/c-api/bool.html?highlight=pybool_type#c.PyBool_Check
func PyBool_Check(o *PyObject) bool {
	return C._go_PyBool_Check(Toc(o)) == 1
}
