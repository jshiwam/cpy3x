package pycore

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import (
	"unsafe"
)

//Long : https://docs.python.org/3/c-api/long.html#c.PyLong_Type
var Long = Togo((*C.PyObject)(unsafe.Pointer(&C.PyLong_Type)))

//PyLong_Check : https://docs.python.org/3/c-api/long.html#c.PyLong_Check
func PyLong_Check(p *PyObject) bool {
	return C._go_PyLong_Check(Toc(p)) != 0
}

//PyLong_CheckExact : https://docs.python.org/3/c-api/long.html#c.PyLong_CheckExact
func PyLong_CheckExact(p *PyObject) bool {
	return C._go_PyLong_CheckExact(Toc(p)) != 0
}

//PyLong_FromLong : https://docs.python.org/3/c-api/long.html#c.PyLong_FromLong
func PyLong_FromLong(v int) *PyObject {
	return Togo(C.PyLong_FromLong(C.long(v)))
}

//PyLong_FromUnsignedLong : https://docs.python.org/3/c-api/long.html#c.PyLong_FromUnsignedLong
func PyLong_FromUnsignedLong(v uint) *PyObject {
	return Togo(C.PyLong_FromUnsignedLong(C.ulong(v)))
}

//PyLong_FromLongLong : https://docs.python.org/3/c-api/long.html#c.PyLong_FromLongLong
func PyLong_FromLongLong(v int64) *PyObject {
	return Togo(C.PyLong_FromLongLong(C.longlong(v)))
}

//PyLong_FromUnsignedLongLong : https://docs.python.org/3/c-api/long.html#c.PyLong_FromUnsignedLongLong
func PyLong_FromUnsignedLongLong(v uint64) *PyObject {
	return Togo(C.PyLong_FromUnsignedLongLong(C.ulonglong(v)))
}

//PyLong_FromDouble : https://docs.python.org/3/c-api/long.html#c.PyLong_FromDouble
func PyLong_FromDouble(v float64) *PyObject {
	return Togo(C.PyLong_FromDouble(C.double(v)))
}

//PyLong_FromString : https://docs.python.org/3/c-api/long.html#c.PyLong_FromString
func PyLong_FromString(str string, base int) *PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return Togo(C.PyLong_FromString(cstr, nil, C.int(base)))
}

//PyLong_FromUnicodeObject : https://docs.python.org/3/c-api/long.html#c.PyLong_FromUnicodeObject
func PyLong_FromUnicodeObject(u *PyObject, base int) *PyObject {
	return Togo(C.PyLong_FromUnicodeObject(Toc(u), C.int(base)))
}

//PyLong_FromGoInt ensures the go integer type does not overflow
func PyLong_FromGoInt(v int) *PyObject {
	return Togo(C.PyLong_FromLongLong(C.longlong(v)))
}

//PyLong_FromGoUint ensures the go integer type does not overflow
func PyLong_FromGoUint(v uint) *PyObject {
	return Togo(C.PyLong_FromUnsignedLongLong(C.ulonglong(v)))
}

//PyLong_FromGoInt64 ensures the go integer type does not overflow
func PyLong_FromGoInt64(v int64) *PyObject {
	return Togo(C.PyLong_FromLongLong(C.longlong(v)))
}

//PyLong_FromGoUint64 ensures the go integer type does not overflow
func PyLong_FromGoUint64(v uint64) *PyObject {
	return Togo(C.PyLong_FromUnsignedLongLong(C.ulonglong(v)))
}

//PyLong_FromGoFloat64 ensures the go integer type does not overflow
func PyLong_FromGoFloat64(v float64) *PyObject {
	return Togo(C.PyLong_FromDouble(C.double(v)))
}

//PyLong_AsLong : https://docs.python.org/3/c-api/long.html#c.PyLong_AsLong
func PyLong_AsLong(obj *PyObject) int {
	return int(C.PyLong_AsLong(Toc(obj)))
}

//PyLong_AsLongAndOverflow : https://docs.python.org/3/c-api/long.html#c.PyLong_AsLongAndOverflow
func PyLong_AsLongAndOverflow(obj *PyObject) (int, int) {
	overflow := C.int(0)
	ret := C.PyLong_AsLongAndOverflow(Toc(obj), &overflow)
	return int(ret), int(overflow)
}

//PyLong_AsLongLong : https://docs.python.org/3/c-api/long.html#c.PyLong_AsLongLong
func PyLong_AsLongLong(obj *PyObject) int64 {
	return int64(C.PyLong_AsLongLong(Toc(obj)))
}

//PyLong_AsLongLongAndOverflow : https://docs.python.org/3/c-api/long.html#c.PyLong_AsLongLongAndOverflow
func PyLong_AsLongLongAndOverflow(obj *PyObject) (int64, int) {
	overflow := C.int(0)
	ret := C.PyLong_AsLongLongAndOverflow(Toc(obj), &overflow)
	return int64(ret), int(overflow)
}

//PyLong_AsUnsignedLong : https://docs.python.org/3/c-api/long.html#c.PyLong_AsUnsignedLong
func PyLong_AsUnsignedLong(obj *PyObject) uint {
	return uint(C.PyLong_AsUnsignedLong(Toc(obj)))
}

//PyLong_AsUnsignedLongLong : https://docs.python.org/3/c-api/long.html#c.PyLong_AsUnsignedLongLong
func PyLong_AsUnsignedLongLong(obj *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongLong(Toc(obj)))
}

//PyLong_AsUnsignedLongMask : https://docs.python.org/3/c-api/long.html#c.PyLong_AsUnsignedLongMask
func PyLong_AsUnsignedLongMask(obj *PyObject) uint {
	return uint(C.PyLong_AsUnsignedLongMask(Toc(obj)))
}

//PyLong_AsUnsignedLongLongMask : https://docs.python.org/3/c-api/long.html#c.PyLong_AsUnsignedLongLongMask
func PyLong_AsUnsignedLongLongMask(obj *PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongLongMask(Toc(obj)))
}

//PyLong_AsDouble : https://docs.python.org/3/c-api/long.html#c.PyLong_AsDouble
func PyLong_AsDouble(obj *PyObject) float64 {
	return float64(C.PyLong_AsDouble(Toc(obj)))
}
