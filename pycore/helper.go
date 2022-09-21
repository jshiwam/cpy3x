package pycore

/*
#include "Python.h"
*/
import "C"
import "unsafe"

// Convert PyObject to C.PyObject
func toc(o *PyObject) *C.PyObject {
	return (*C.PyObject)(o)
}

// Convert C.PyObject to PyObject
func togo(o *C.PyObject) *PyObject {
	return (*PyObject)(o)
}

func VoidToC(o unsafe.Pointer) *C.PyObject {
	return (*C.PyObject)(o)
}

// Convert C.PyObject to PyObject
func Togo(o *C.PyObject) *PyObject {
	return (*PyObject)(o)
}
