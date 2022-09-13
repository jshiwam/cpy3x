package pycore

// #cgo pkg-config: python3
// #include "Python.h"
import "C"

type PyObject C.PyObject

var (
	Py_False = C.Py_False
	Py_True  = C.Py_True
)

func PyBool_FromLong(v int) *PyObject {
	return (*PyObject)(C.PyBool_FromLong(C.long(v)))
}
