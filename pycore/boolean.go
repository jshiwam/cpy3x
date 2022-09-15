package pycore

// #cgo pkg-config: python35
// #include "Python.h"
import "C"

type PyObject C.PyObject

var (
	Py_False = (*PyObject)(C.Py_False)
	Py_True  = (*PyObject)(C.Py_True)
)

func PyBool_FromLong(v int) *PyObject {
	return (*PyObject)(C.PyBool_FromLong(C.long(v)))
}
