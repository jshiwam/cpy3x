package pycore

/*
#include "Python.h"
#include "macros.h"
*/
import "C"
import "unsafe"

var ByteArray = (*PyObject)((*C.PyObject)(unsafe.Pointer(&C.PyByteArray_Type)))

func PyByteArray_Check(o *PyObject) bool {
	return C._go_PyByteArray_Check((*C.PyObject)(o)) == 1
}
