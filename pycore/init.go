package pycore

// #include "Python.h"
import "C"

func Py_Initialize() {
	C.Py_Initialize()
}

func Py_Finalize() {
	C.Py_Finalize()
}
