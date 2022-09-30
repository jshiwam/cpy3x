package pycore

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

//Float : https://docs.python.org/3/c-api/float.html#c.PyFloat_Type
var Float = Togo((*C.PyObject)(unsafe.Pointer(&C.PyFloat_Type)))

//PyFloat_Check : https://docs.python.org/3/c-api/float.html#c.PyFloat_Check
func PyFloat_Check(p *PyObject) bool {
	return C._go_PyFloat_Check(Toc(p)) != 0
}

//PyFloat_CheckExact : https://docs.python.org/3/c-api/float.html#c.PyFloat_CheckExact
func PyFloat_CheckExact(p *PyObject) bool {
	return C._go_PyFloat_CheckExact(Toc(p)) != 0
}

//PyFloat_FromDouble : https://docs.python.org/3/c-api/float.html#c.PyFloat_FromDouble
func PyFloat_FromDouble(v float64) *PyObject {
	return Togo(C.PyFloat_FromDouble(C.double(v)))
}

//PyFloat_FromString : https://docs.python.org/3/c-api/float.html#c.PyFloat_FromString
func PyFloat_FromString(str *PyObject) *PyObject {
	return Togo(C.PyFloat_FromString(Toc(str)))
}

//PyFloat_AsDouble : https://docs.python.org/3/c-api/float.html#c.PyFloat_AsDouble
func PyFloat_AsDouble(obj *PyObject) float64 {
	return float64(C.PyFloat_AsDouble(Toc(obj)))
}

//PyFloat_GetInfo : https://docs.python.org/3/c-api/float.html#c.PyFloat_GetInfo
func PyFloat_GetInfo() *PyObject {
	return Togo(C.PyFloat_GetInfo())
}

//PyFloat_GetMax : https://docs.python.org/3/c-api/float.html#c.PyFloat_GetMax
func PyFloat_GetMax() float64 {
	return float64(C.PyFloat_GetMax())
}

//PyFloat_GetMin : https://docs.python.org/3/c-api/float.html#c.PyFloat_GetMin
func PyFloat_GetMin() float64 {
	return float64(C.PyFloat_GetMin())
}

// Change : Removed 3.9 onwards
//PyFloat_ClearFreeList : https://docs.python.org/3/c-api/float.html#c.PyFloat_ClearFreeList
// func PyFloat_ClearFreeList() int {
// 	return int(C.PyFloat_ClearFreeList())
// }
