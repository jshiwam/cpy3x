package pycore

//go:generate go run script/variadic.go

/*
#include "Python.h"
#include "macro.h"
#include "variadic.h"
*/
import "C"
import (
	"unsafe"
)

//MaxVariadicLength is the maximum number of arguments that can be passed to a variadic C function due to a cgo limitation
var MaxVariadicLength = 20

// Constants used for comparison in PyObject_RichCompareBool
var (
	Py_LT = int(C.Py_LT)
	Py_LE = int(C.Py_LE)
	Py_EQ = int(C.Py_EQ)
	Py_NE = int(C.Py_NE)
	Py_GT = int(C.Py_GT)
	Py_GE = int(C.Py_GE)
)

//Not a part of stable ABI but present across versions 3.5-3.10
//None : https://docs.python.org/3/c-api/none.html#c.Py_None
var Py_None = Togo(C.Py_None)

//Part of Stable ABI
//PyObject : https://docs.python.org/3/c-api/structures.html?highlight=pyobject#c.PyObject
type PyObject C.PyObject

//Part of Stable ABI
//IncRef : https://docs.python.org/3/c-api/refcounting.html#c.Py_INCREF
func (pyObject *PyObject) IncRef() {
	C.Py_IncRef(Toc(pyObject))
}

//Part of Stable ABI
//DecRef : https://docs.python.org/3/c-api/refcounting.html#c.Py_DECREF
func (pyObject *PyObject) DecRef() {
	C.Py_DecRef(Toc(pyObject))
}

//Part of Stable ABI
//ReprEnter : https://docs.python.org/3/c-api/exceptions.html#c.Py_ReprEnter
func (pyObject *PyObject) ReprEnter() int {
	return int(C.Py_ReprEnter(Toc(pyObject)))
}

//Part of Stable ABI
//ReprLeave : https://docs.python.org/3/c-api/exceptions.html#c.Py_ReprLeave
func (pyObject *PyObject) ReprLeave() {
	C.Py_ReprLeave(Toc(pyObject))
}

//Part of Stable ABI
//HasAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_HasAttr
func (pyObject *PyObject) HasAttr(attr_name *PyObject) bool {
	return C.PyObject_HasAttr(Toc(pyObject), Toc(attr_name)) == 1
}

//Part of Stable ABI
//HasAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_HasAttr
func (pyObject *PyObject) HasAttrString(attr_name string) bool {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return C.PyObject_HasAttrString(Toc(pyObject), cattr_name) == 1
}

//Part of Stable ABI
//GetAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_GetAttr
func (pyObject *PyObject) GetAttr(attr_name *PyObject) *PyObject {
	return Togo(C.PyObject_GetAttr(Toc(pyObject), Toc(attr_name)))
}

//Part of Stable ABI
//GetAttrString : https://docs.python.org/3/c-api/object.html#c.PyObject_GetAttrString
func (pyObject *PyObject) GetAttrString(attr_name string) *PyObject {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return Togo(C.PyObject_GetAttrString(Toc(pyObject), cattr_name))
}

//Part of Stable ABI
//SetAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_SetAttr
func (pyObject *PyObject) SetAttr(attr_name *PyObject, v *PyObject) int {
	return int(C.PyObject_SetAttr(Toc(pyObject), Toc(attr_name), Toc(v)))
}

//Part of Stable ABI
//SetAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_SetAttr
func (pyObject *PyObject) SetAttrString(attr_name string, v *PyObject) int {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return int(C.PyObject_SetAttrString(Toc(pyObject), cattr_name, Toc(v)))
}

//Not a part of stable ABI but present across versions 3.5-3.10
//DelAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_DelAttr
func (pyObject *PyObject) DelAttr(attr_name *PyObject) int {
	return int(C._go_PyObject_DelAttr(Toc(pyObject), Toc(attr_name)))
}

//Not a part of stable ABI but present across versions 3.5-3.10
//DelAttrString : https://docs.python.org/3/c-api/object.html#c.PyObject_DelAttrString
func (pyObject *PyObject) DelAttrString(attr_name string) int {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return int(C._go_PyObject_DelAttrString(Toc(pyObject), cattr_name))
}

//Part of Stable ABI
//RichCompare : https://docs.python.org/3/c-api/object.html#c.PyObject_RichCompare
func (pyObject *PyObject) RichCompare(o *PyObject, opid int) *PyObject {
	return Togo(C.PyObject_RichCompare(Toc(pyObject), Toc(o), C.int(opid)))
}

//Part of Stable ABI
//RichCompareBool : https://docs.python.org/3/c-api/object.html#c.PyObject_RichCompareBool
func (pyObject *PyObject) RichCompareBool(o *PyObject, opid int) int {
	return int(C.PyObject_RichCompareBool(Toc(pyObject), Toc(o), C.int(opid)))
}

//Part of Stable ABI
//Repr : https://docs.python.org/3/c-api/object.html#c.PyObject_Repr
func (pyObject *PyObject) Repr() *PyObject {
	return Togo(C.PyObject_Repr(Toc(pyObject)))
}

//Part of Stable ABI
//ASCII : https://docs.python.org/3/c-api/object.html#c.PyObject_ASCII
func (pyObject *PyObject) ASCII() *PyObject {
	return Togo(C.PyObject_ASCII(Toc(pyObject)))
}

//Part of Stable ABI
//Str : https://docs.python.org/3/c-api/object.html#c.PyObject_Str
func (pyObject *PyObject) Str() *PyObject {
	return Togo(C.PyObject_Str(Toc(pyObject)))
}

//Part of Stable ABI
//Bytes :  https://docs.python.org/3/c-api/object.html#c.PyObject_Bytes
func (pyObject *PyObject) Bytes() *PyObject {
	return Togo(C.PyObject_Bytes(Toc(pyObject)))
}

//Part of Stable ABI
//IsSubclass : https://docs.python.org/3/c-api/object.html#c.PyObject_IsSubclass
func (pyObject *PyObject) IsSubclass(cls *PyObject) int {
	return int(C.PyObject_IsSubclass(Toc(pyObject), Toc(cls)))
}

//Part of Stable ABI
//IsInstance : https://docs.python.org/3/c-api/object.html#c.PyObject_IsInstance
func (pyObject *PyObject) IsInstance(cls *PyObject) int {
	return int(C.PyObject_IsInstance(Toc(pyObject), Toc(cls)))
}

//Part of Stable ABI
//PyCallable_Check : https://docs.python.org/3/c-api/call.html#c.PyCallable_Check
func PyCallable_Check(pyObject *PyObject) bool {
	return C.PyCallable_Check(Toc(pyObject)) == 1
}

//Part of Stable ABI
//Call : https://docs.python.org/3/c-api/call.html#c.PyObject_Call
func (pyObject *PyObject) Call(args *PyObject, kwargs *PyObject) *PyObject {
	return Togo(C.PyObject_Call(Toc(pyObject), Toc(args), Toc(kwargs)))
}

//Part of Stable ABI
//CallObject : https://docs.python.org/3/c-api/call.html#c.PyObject_CallObject
func (pyObject *PyObject) CallObject(args *PyObject) *PyObject {
	return Togo(C.PyObject_CallObject(Toc(pyObject), Toc(args)))
}

//Part of Stable ABI
//CallFunctionObjArgs : https://docs.python.org/3/c-api/call.html#c.PyObject_CallFunctionObjArgs
func (pyObject *PyObject) CallFunctionObjArgs(args ...*PyObject) *PyObject {
	if len(args) > MaxVariadicLength {
		panic("CallFunctionObjArgs: too many arguments")
	}
	if len(args) == 0 {
		return Togo(C._go_PyObject_CallFunctionObjArgs(Toc(pyObject), 0, (**C.PyObject)(nil)))
	}
	cargs := make([]*C.PyObject, len(args), len(args))
	for i, arg := range args {
		cargs[i] = Toc(arg)
	}
	return Togo(C._go_PyObject_CallFunctionObjArgs(Toc(pyObject), C.int(len(args)), (**C.PyObject)(unsafe.Pointer(&cargs[0]))))
}

//Part of Stable ABI
//CallMethodObjArgs : https://docs.python.org/3/c-api/call.html#c.PyObject_CallMethodObjArgs
func (pyObject *PyObject) CallMethodObjArgs(name *PyObject, args ...*PyObject) *PyObject {

	if len(args) > MaxVariadicLength {
		panic("CallMethodObjArgs: too many arguments ")
	}

	if len(args) == 0 {
		return Togo(C._go_PyObject_CallMethodObjArgs(Toc(pyObject), Toc(name), 0, (**C.PyObject)(nil)))
	}
	cargs := make([]*C.PyObject, len(args), len(args))
	for i, arg := range args {
		cargs[i] = Toc(arg)
	}
	return Togo(C._go_PyObject_CallMethodObjArgs(Toc(pyObject), Toc(name), C.int(len(args)), (**C.PyObject)(unsafe.Pointer(&cargs[0]))))
}

//CallMethodArgs : same as CallMethodObjArgs but with name as go string
func (pyObject *PyObject) CallMethodArgs(name string, args ...*PyObject) *PyObject {
	nameObj := PyUnicode_FromString(name)
	defer nameObj.DecRef()

	return pyObject.CallMethodObjArgs(nameObj, args...)
}

//Part of Stable ABI
//Hash : https://docs.python.org/3/c-api/object.html#c.PyObject_Hash
func (pyObject *PyObject) Hash() int {
	return int(C.PyObject_Hash(Toc(pyObject)))
}

//Part of Stable ABI
//HashNotImplemented : https://docs.python.org/3/c-api/object.html#c.PyObject_HashNotImplemented
func (pyObject *PyObject) HashNotImplemented() int {
	return int(C.PyObject_HashNotImplemented(Toc(pyObject)))
}

//Part of Stable ABI
//IsTrue : https://docs.python.org/3/c-api/object.html#c.PyObject_IsTrue
func (pyObject *PyObject) IsTrue() int {
	return int(C.PyObject_IsTrue(Toc(pyObject)))
}

//Part of Stable ABI
//Not : https://docs.python.org/3/c-api/object.html#c.PyObject_Not
func (pyObject *PyObject) Not() int {
	return int(C.PyObject_Not(Toc(pyObject)))
}

//Part of Stable ABI
//Type : https://docs.python.org/3/c-api/object.html#c.PyObject_Type
func (pyObject *PyObject) Type() *PyObject {
	return Togo(C.PyObject_Type(Toc(pyObject)))
}

//Part of Stable ABI
//Length : https://docs.python.org/3/c-api/object.html#c.PyObject_Length
func (pyObject *PyObject) Length() int {
	return int(C.PyObject_Length(Toc(pyObject)))
}

//Not a part of stable ABI but present across versions 3.5-3.10
//LengthHint : https://docs.python.org/3.5/c-api/object.html#c.PyObject_LengthHint
func (pyObject *PyObject) LengthHint(pyDefault int) int {
	return int(C.PyObject_LengthHint(Toc(pyObject), C.Py_ssize_t(pyDefault)))
}

//Part of Stable ABI
//GetItem : https://docs.python.org/3/c-api/object.html#c.PyObject_GetItem
func (pyObject *PyObject) GetItem(key *PyObject) *PyObject {
	return Togo(C.PyObject_GetItem(Toc(pyObject), Toc(key)))
}

//Part of Stable ABI
//SetItem : https://docs.python.org/3/c-api/object.html#c.PyObject_SetItem
func (pyObject *PyObject) SetItem(key *PyObject, val *PyObject) int {
	return int(C.PyObject_SetItem(Toc(pyObject), Toc(key), Toc(val)))
}

//Not a part of stable ABI but present across versions 3.5-3.10
//DelItem : https://docs.python.org/3/c-api/object.html#c.PyObject_DelItem
func (pyObject *PyObject) DelItem(key *PyObject) int {
	return int(C.PyObject_DelItem(Toc(pyObject), Toc(key)))
}

//Part of Stable ABI
//Dir : https://docs.python.org/3/c-api/object.html#c.PyObject_Dir
func (pyObject *PyObject) Dir() *PyObject {
	return Togo(C.PyObject_Dir(Toc(pyObject)))
}

//Part of Stable ABI
//GetIter : https://docs.python.org/3/c-api/object.html#c.PyObject_GetIter
func (pyObject *PyObject) GetIter() *PyObject {
	return Togo(C.PyObject_GetIter(Toc(pyObject)))
}

// // Change: Part of Stable ABI Should Go In Version 3.10
// //GetAIter : https://docs.python.org/3/c-api/object.html#c.PyObject_GetAIter
// func (pyObject *PyObject) GetAIter() *PyObject {
// 	return Togo(C.PyObject_GetAIter(Toc(pyObject)))
// }
