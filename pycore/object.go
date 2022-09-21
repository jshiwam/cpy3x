package pycore

import "unsafe"

/*
#include "Python.h"
#include "macro.h"
*/

//Not a part of stable ABI but present across versions 3.5-3.10
//None : https://docs.python.org/3/c-api/none.html#c.Py_None
var Py_None = togo(C.Py_None)

//Part of Stable ABI
//PyObject : https://docs.python.org/3/c-api/structures.html?highlight=pyobject#c.PyObject
type PyObject C.PyObject

//Part of Stable ABI
//IncRef : https://docs.python.org/3/c-api/refcounting.html#c.Py_INCREF
func (pyObject *PyObject) IncRef() {
	C.Py_IncRef(toc(pyObject))
}

//Part of Stable ABI
//DecRef : https://docs.python.org/3/c-api/refcounting.html#c.Py_DECREF
func (pyObject *PyObject) DecRef() {
	C.Py_DecRef(toc(pyObject))
}

//Part of Stable ABI
//ReprEnter : https://docs.python.org/3/c-api/exceptions.html#c.Py_ReprEnter
func (pyObject *PyObject) ReprEnter() int {
	return int(C.Py_ReprEnter(toc(pyObject)))
}

//Part of Stable ABI
//ReprLeave : https://docs.python.org/3/c-api/exceptions.html#c.Py_ReprLeave
func (pyObject *PyObject) ReprLeave() {
	C.Py_ReprLeave(toc(pyObject))
}

//Part of Stable ABI
//HasAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_HasAttr
func (pyObject *PyObject) HasAttr(attr_name *PyObject) bool {
	return C.PyObject_HasAttr(toc(pyObject), toc(attr_name)) == 1
}

//Part of Stable ABI
//HasAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_HasAttr
func (pyObject *PyObject) HasAttrString(attr_name string) bool {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return C.PyObject_HasAttrString(toc(pyObject), cattr_name) == 1
}

//Part of Stable ABI
//GetAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_GetAttr
func (pyObject *PyObject) GetAttr(attr_name *PyObject) *PyObject {
	return togo(C.PyObject_GetAttr(toc(pyObject), toc(attr_name)))
}

//Part of Stable ABI
//GetAttrString : https://docs.python.org/3/c-api/object.html#c.PyObject_GetAttrString
func (pyObject *PyObject) GetAttrString(attr_name string) *PyObject {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return togo(C.PyObject_GetAttrString(toc(pyObject), cattr_name))
}

//Part of Stable ABI
//SetAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_SetAttr
func (pyObject *PyObject) SetAttr(attr_name *PyObject, v *PyObject) int {
	return int(C.PyObject_SetAttr(toc(pyObject), toc(attr_name), toc(v)))
}

//Part of Stable ABI
//SetAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_SetAttr
func (pyObject *PyObject) SetAttrString(attr_name string, v *PyObject) int {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return int(C.PyObject_SetAttrString(toc(pyObject), cattr_name, toc(v)))
}

//Not a part of stable ABI but present across versions 3.5-3.10
//DelAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_DelAttr
func (pyObject *PyObject) DelAttr(attr_name *PyObject) int {
	return int(C._go_PyObject_DelAttr(toc(pyObject), toc(attr_name)))
}

//Not a part of stable ABI but present across versions 3.5-3.10
//DelAttrString : https://docs.python.org/3/c-api/object.html#c.PyObject_DelAttrString
func (pyObject *PyObject) DelAttrString(attr_name string) int {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return int(C._go_PyObject_DelAttrString(toc(pyObject), cattr_name))
}

//Part of Stable ABI
//RichCompare : https://docs.python.org/3/c-api/object.html#c.PyObject_RichCompare
func (pyObject *PyObject) RichCompare(o *PyObject, opid int) *PyObject {
	return togo(C.PyObject_RichCompare(toc(pyObject), toc(o), C.int(opid)))
}

//Part of Stable ABI
//RichCompareBool : https://docs.python.org/3/c-api/object.html#c.PyObject_RichCompareBool
func (pyObject *PyObject) RichCompareBool(o *PyObject, opid int) int {
	return int(C.PyObject_RichCompareBool(toc(pyObject), toc(o), C.int(opid)))
}
