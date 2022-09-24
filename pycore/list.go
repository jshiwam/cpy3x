package pycore

/*
#include "Python.h"
*/
import "C"

//PyList_New : https://docs.python.org/3/c-api/list.html#c.PyList_New
func PyList_New(len int) *PyObject {
	return Togo(C.PyList_New(C.Py_ssize_t(len)))
}
