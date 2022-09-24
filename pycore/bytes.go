package pycore

/*
#include "Python.h"
*/
import "C"
import "unsafe"

func PyBytes_FromString(str string) *PyObject {
	cstring := C.CString(str)
	defer C.free(unsafe.Pointer(cstring))

	return Togo(C.PyBytes_FromString(cstring))
}
