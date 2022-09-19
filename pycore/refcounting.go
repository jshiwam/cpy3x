package pycore

/*
#include "Python.h"
*/
import "C"

func DecRef(o *PyObject) {
	C.Py_DecRef((*C.PyObject)(o))
}

func IncRef(o *PyObject) {
	C.Py_IncRef((*C.PyObject)(o))
}
