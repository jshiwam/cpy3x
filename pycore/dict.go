package pycore

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import (
	"unsafe"
)

//Dict : https://docs.python.org/3/c-api/dict.html#c.PyDict_Type
var Dict = Togo((*C.PyObject)(unsafe.Pointer(&C.PyDict_Type)))

//PyDict_Check : https://docs.python.org/3/c-api/dict.html#c.PyDict_Check
func PyDict_Check(p *PyObject) bool {
	return C._go_PyDict_Check(Toc(p)) != 0
}

//PyDict_CheckExact : https://docs.python.org/3/c-api/dict.html#c.PyDict_CheckExact
func PyDict_CheckExact(p *PyObject) bool {
	return C._go_PyDict_CheckExact(Toc(p)) != 0
}

//PyDict_New : https://docs.python.org/3/c-api/dict.html#c.PyDict_New
func PyDict_New() *PyObject {
	return Togo(C.PyDict_New())
}

//PyDictProxy_New : https://docs.python.org/3/c-api/dict.html#c.PyDictProxy_New
func PyDictProxy_New(mapping *PyObject) *PyObject {
	return Togo(C.PyDictProxy_New(Toc(mapping)))
}

//PyDict_Clear : https://docs.python.org/3/c-api/dict.html#c.PyDict_Clear
func PyDict_Clear(p *PyObject) {
	C.PyDict_Clear(Toc(p))
}

//PyDict_Contains : https://docs.python.org/3/c-api/dict.html#c.PyDict_Contains
func PyDict_Contains(p, key *PyObject) int {
	return int(C.PyDict_Contains(Toc(p), Toc(key)))
}

//PyDict_Copy : https://docs.python.org/3/c-api/dict.html#c.PyDict_Copy
func PyDict_Copy(p *PyObject) *PyObject {
	return Togo(C.PyDict_Copy(Toc(p)))
}

//PyDict_SetItem : https://docs.python.org/3/c-api/dict.html#c.PyDict_SetItem
func PyDict_SetItem(p, key, val *PyObject) int {
	return int(C.PyDict_SetItem(Toc(p), Toc(key), Toc(val)))
}

//PyDict_SetItemString : https://docs.python.org/3/c-api/dict.html#c.PyDict_SetItemString
func PyDict_SetItemString(p *PyObject, key string, val *PyObject) int {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	return int(C.PyDict_SetItemString(Toc(p), ckey, Toc(val)))
}

//PyDict_DelItem : https://docs.python.org/3/c-api/dict.html#c.PyDict_DelItem
func PyDict_DelItem(p, key *PyObject) int {
	return int(C.PyDict_DelItem(Toc(p), Toc(key)))
}

//PyDict_DelItemString : https://docs.python.org/3/c-api/dict.html#c.PyDict_DelItemString
func PyDict_DelItemString(p *PyObject, key string) int {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	return int(C.PyDict_DelItemString(Toc(p), ckey))
}

//PyDict_GetItem : https://docs.python.org/3/c-api/dict.html#c.PyDict_GetItem
func PyDict_GetItem(p, key *PyObject) *PyObject {
	return Togo(C.PyDict_GetItem(Toc(p), Toc(key)))
}

//PyDict_GetItemWithError : https://docs.python.org/3/c-api/dict.html#c.PyDict_GetItemWithError
func PyDict_GetItemWithError(p, key *PyObject) *PyObject {
	return Togo(C.PyDict_GetItemWithError(Toc(p), Toc(key)))
}

//PyDict_GetItemString : https://docs.python.org/3/c-api/dict.html#c.PyDict_GetItemString
func PyDict_GetItemString(p *PyObject, key string) *PyObject {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	return Togo(C.PyDict_GetItemString(Toc(p), ckey))
}

//PyDict_SetDefault : https://docs.python.org/3/c-api/dict.html#c.PyDict_SetDefault
func PyDict_SetDefault(p, key, pyDefault *PyObject) *PyObject {
	return Togo(C.PyDict_SetDefault(Toc(p), Toc(key), Toc(pyDefault)))
}

//PyDict_Items : https://docs.python.org/3/c-api/dict.html#c.PyDict_Items
func PyDict_Items(p *PyObject) *PyObject {
	return Togo(C.PyDict_Items(Toc(p)))
}

//PyDict_Keys : https://docs.python.org/3/c-api/dict.html#c.PyDict_Keys
func PyDict_Keys(p *PyObject) *PyObject {
	return Togo(C.PyDict_Keys(Toc(p)))
}

//PyDict_Values : https://docs.python.org/3/c-api/dict.html#c.PyDict_Values
func PyDict_Values(p *PyObject) *PyObject {
	return Togo(C.PyDict_Values(Toc(p)))
}

//PyDict_Size : https://docs.python.org/3/c-api/dict.html#c.PyDict_Size
func PyDict_Size(p *PyObject) int {
	return int(C.PyDict_Size(Toc(p)))
}

//PyDict_Next : https://docs.python.org/3/c-api/dict.html#c.PyDict_Next
func PyDict_Next(p *PyObject, ppos *int, pkey, pvalue **PyObject) bool {
	cpos := C.Py_ssize_t(*ppos)
	ckey := Toc(*pkey)
	cvalue := Toc(*pvalue)

	res := C.PyDict_Next(Toc(p), &cpos, &ckey, &cvalue) != 0

	*ppos = int(cpos)
	*pkey = Togo(ckey)
	*pvalue = Togo(cvalue)

	return res
}

// //PyDict_ClearFreeList : https://docs.python.org/3/c-api/dict.html#c.PyDict_ClearFreeList
// func PyDict_ClearFreeList() int {
// 	return int(C.PyDict_ClearFreeList())
// }
