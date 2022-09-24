package pycore

/*
#include "Python.h"
#include "macros.h"
*/
import "C"
import "unsafe"

//PyUnicode_FromString : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_FromString
func PyUnicode_FromString(u string) *PyObject {
	cu := C.CString(u)
	defer C.free(unsafe.Pointer(cu))

	return Togo(C.PyUnicode_FromString(cu))
}

//PyUnicode_AsUTF8 : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_AsUTF8
func PyUnicode_AsUTF8(unicode *PyObject) string {
	cutf8 := C.PyUnicode_AsUTF8(Toc(unicode))
	return C.GoString(cutf8)
}
