package pycore

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

//Type : https://docs.python.org/3/c-api/type.html#c.PyType_Type
var Type = Togo((*C.PyObject)(unsafe.Pointer(&C.PyType_Type)))

//PyType_Check : https://docs.python.org/3/c-api/type.html#c.PyType_Check
func PyType_Check(o *PyObject) bool {
	return C._go_PyType_Check(Toc(o)) != 0
}

//PyType_CheckExact : https://docs.python.org/3/c-api/type.html#c.PyType_CheckExact
func PyType_CheckExact(o *PyObject) bool {
	return C._go_PyType_CheckExact(Toc(o)) != 0
}
