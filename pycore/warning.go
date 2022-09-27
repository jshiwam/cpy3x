package pycore

/*
#include "Python.h"
*/
import "C"

import (
	"unsafe"
)

/*
All standard Python warning categories are available as global variables whose names are PyExc_ followed by the Python exception name.
These have the type PyObject*; they are all class objects.
*/
var (
	PyExc_Warning                   = Togo(C.PyExc_Warning)
	PyExc_BytesWarning              = Togo(C.PyExc_BytesWarning)
	PyExc_DeprecationWarning        = Togo(C.PyExc_DeprecationWarning)
	PyExc_FutureWarning             = Togo(C.PyExc_FutureWarning)
	PyExc_ImportWarning             = Togo(C.PyExc_ImportWarning)
	PyExc_PendingDeprecationWarning = Togo(C.PyExc_PendingDeprecationWarning)
	PyExc_ResourceWarning           = Togo(C.PyExc_ResourceWarning)
	PyExc_RuntimeWarning            = Togo(C.PyExc_RuntimeWarning)
	PyExc_SyntaxWarning             = Togo(C.PyExc_SyntaxWarning)
	PyExc_UnicodeWarning            = Togo(C.PyExc_UnicodeWarning)
	PyExc_UserWarning               = Togo(C.PyExc_UserWarning)
)

//PyErr_WarnEx : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnEx
func PyErr_WarnEx(category *PyObject, message string, stack_level int) int {
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))

	return int(C.PyErr_WarnEx(Toc(category), cmessage, C.Py_ssize_t(stack_level)))
}

//PyErr_WarnExplicitObject : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnExplicitObject
func PyErr_WarnExplicitObject(category *PyObject, message *PyObject, filename *PyObject, lineno int, module *PyObject, registry *PyObject) int {
	return int(C.PyErr_WarnExplicitObject(Toc(category), Toc(message), Toc(filename), C.int(lineno), Toc(module), Toc(registry)))
}

//PyErr_WarnExplicit : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnExplicit
func PyErr_WarnExplicit(category *PyObject, message string, filename string, lineno int, module string, registry *PyObject) int {
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	cmodule := C.CString(module)
	defer C.free(unsafe.Pointer(cmodule))

	return int(C.PyErr_WarnExplicit(Toc(category), cmessage, cfilename, C.int(lineno), cmodule, Toc(registry)))
}
