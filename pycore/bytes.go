package pycore

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

//Bytes : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Type
var Bytes = Togo((*C.PyObject)(unsafe.Pointer(&C.PyBytes_Type)))

//PyBytes_Check : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Check
func PyBytes_Check(o *PyObject) bool {
	return C._go_PyBytes_Check(Toc(o)) != 0
}

//PyBytes_CheckExact : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_CheckExact
func PyBytes_CheckExact(o *PyObject) bool {
	return C._go_PyBytes_CheckExact(Toc(o)) != 0
}

//PyBytes_FromString : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_FromString
func PyBytes_FromString(str string) *PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return Togo(C.PyBytes_FromString(cstr))
}

//PyBytes_FromObject : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_FromObject
func PyBytes_FromObject(o *PyObject) *PyObject {
	return Togo(C.PyBytes_FromObject(Toc(o)))
}

//PyBytes_Size : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Size
func PyBytes_Size(o *PyObject) int {
	return int(C.PyBytes_Size(Toc(o)))
}

//PyBytes_AsString : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_AsString
func PyBytes_AsString(o *PyObject) string {
	return C.GoStringN(C.PyBytes_AsString(Toc(o)), C.int(C.PyBytes_Size(Toc(o))))
}

//PyBytes_Concat : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Concat
func PyBytes_Concat(bytes, newpart *PyObject) *PyObject {
	cbytes := Toc(bytes)
	C.PyBytes_Concat(&cbytes, Toc(newpart))
	return Togo(cbytes)
}

//PyBytes_ConcatAndDel : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_ConcatAndDel
func PyBytes_ConcatAndDel(bytes, newpart *PyObject) *PyObject {
	cbytes := Toc(bytes)
	C.PyBytes_ConcatAndDel(&cbytes, Toc(newpart))
	return Togo(cbytes)
}
