package pycore

/*
#include "Python.h"
#include "macros.h"
*/
import "C"
import (
	"unsafe"
)

//ByteArray : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Type
var ByteArray = Togo((*C.PyObject)(unsafe.Pointer(&C.PyByteArray_Type)))

// Not a part of Stable ABI
// PyByteArray_Check : https://docs.python.org/3.10/c-api/bytearray.html#c.PyByteArray_Check
func PyByteArray_Check(o *PyObject) bool {
	return C._go_PyByteArray_Check(Toc(o)) == 1
}

// Not a part of Stable ABI
// PyByteArray_CheckExact: https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_CheckExact
func PyByteArray_CheckExact(o *PyObject) bool {
	return C._go_PyByteArray_CheckExact(Toc(o)) == 1
}

// Part of Stable ABI
// PyByteArray_FromStringAndSize : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_FromStringAndSize
func PyByteArray_FromStringAndSize(str string) *PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return Togo(C.PyByteArray_FromStringAndSize(cstr, C.Py_ssize_t(len(str))))
}

// Part of Stable ABI
// PyByteArray_AsString : https://docs.python.org/3.10/c-api/bytearray.html#c.PyByteArray_AsString
func PyByteArray_AsString(o *PyObject) string {
	return C.GoStringN(C.PyByteArray_AsString(Toc(o)), C.int(PyByteArray_Size(o)))
}

// Part of Stable ABI
// PyByteArray_Resize : https://docs.python.org/3.10/c-api/bytearray.html#c.PyByteArray_Resize
func PyByteArray_Resize(o *PyObject, length int) {
	C.PyByteArray_Resize(Toc(o), C.Py_ssize_t(length))
}

// Part of Stable ABI
// PyByteArray_Size : https://docs.python.org/3.10/c-api/bytearray.html#c.PyByteArray_Size
func PyByteArray_Size(o *PyObject) int {
	return int(C.PyByteArray_Size(Toc(o)))
}

// Part of Stable ABI
// PyByteArray_FromObject : https://docs.python.org/3.10/c-api/bytearray.html#c.PyByteArray_FromObject
func PyByteArray_FromObject(o *PyObject) *PyObject {
	return Togo(C.PyByteArray_FromObject(Toc(o)))
}

// Part of Stable ABI
// PyByteArray_Concat : https://docs.python.org/3.10/c-api/bytearray.html#c.PyByteArray_Concat
func PyByteArray_Concat(a *PyObject, b *PyObject) *PyObject {
	return Togo(C.PyByteArray_Concat(Toc(a), Toc(b)))
}

func pythonRepr(o *PyObject) string {
	return C.GoString(C.PyUnicode_AsUTF8(C.PyObject_Repr(Toc(o))))
}
