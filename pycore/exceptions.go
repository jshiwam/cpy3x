package pycore

/*
#include "Python.h"
*/
import "C"

import (
	"unsafe"
)

/*
All standard Python exceptions are available as global variables whose names are PyExc_ followed by the Python exception name.
These have the type PyObject*; they are all class objects.
*/
var (
	PyExc_BaseException          = Togo(C.PyExc_BaseException)
	PyExc_Exception              = Togo(C.PyExc_Exception)
	PyExc_ArithmeticError        = Togo(C.PyExc_ArithmeticError)
	PyExc_AssertionError         = Togo(C.PyExc_AssertionError)
	PyExc_AttributeError         = Togo(C.PyExc_AttributeError)
	PyExc_BlockingIOError        = Togo(C.PyExc_BlockingIOError)
	PyExc_BrokenPipeError        = Togo(C.PyExc_BrokenPipeError)
	PyExc_BufferError            = Togo(C.PyExc_BufferError)
	PyExc_ChildProcessError      = Togo(C.PyExc_ChildProcessError)
	PyExc_ConnectionAbortedError = Togo(C.PyExc_ConnectionAbortedError)
	PyExc_ConnectionError        = Togo(C.PyExc_ConnectionError)
	PyExc_ConnectionRefusedError = Togo(C.PyExc_ConnectionRefusedError)
	PyExc_ConnectcionResetError  = Togo(C.PyExc_ConnectionResetError)
	PyExc_EOFError               = Togo(C.PyExc_EOFError)
	PyExc_FileExistsError        = Togo(C.PyExc_FileExistsError)
	PyExc_FileNotFoundError      = Togo(C.PyExc_FileNotFoundError)
	PyExc_FloatingPointError     = Togo(C.PyExc_FloatingPointError)
	PyExc_GeneratorExit          = Togo(C.PyExc_GeneratorExit)
	PyExc_ImportError            = Togo(C.PyExc_ImportError)
	PyExc_IndentationError       = Togo(C.PyExc_IndentationError)
	PyExc_IndexError             = Togo(C.PyExc_IndexError)
	PyExc_InterruptedError       = Togo(C.PyExc_InterruptedError)
	PyExc_IsADirectoryError      = Togo(C.PyExc_IsADirectoryError)
	PyExc_KeyError               = Togo(C.PyExc_KeyError)
	PyExc_KeyboardInterrupt      = Togo(C.PyExc_KeyboardInterrupt)
	PyExc_LookupError            = Togo(C.PyExc_LookupError)
	PyExc_MemoryError            = Togo(C.PyExc_MemoryError)
	// Change: New in 3.6
	// PyExc_ModuleNotFoundError    = Togo(C.PyExc_ModuleNotFoundError)
	PyExc_NameError             = Togo(C.PyExc_NameError)
	PyExc_NotADirectoryError    = Togo(C.PyExc_NotADirectoryError)
	PyExc_NotImplementedError   = Togo(C.PyExc_NotImplementedError)
	PyExc_OSError               = Togo(C.PyExc_OSError)
	PyExc_OverflowError         = Togo(C.PyExc_OverflowError)
	PyExc_PermissionError       = Togo(C.PyExc_PermissionError)
	PyExc_ProcessLookupError    = Togo(C.PyExc_ProcessLookupError)
	PyExc_RecursionError        = Togo(C.PyExc_RecursionError)
	PyExc_ReferenceError        = Togo(C.PyExc_ReferenceError)
	PyExc_RuntimeError          = Togo(C.PyExc_RuntimeError)
	PyExc_StopAsyncIteration    = Togo(C.PyExc_StopAsyncIteration)
	PyExc_StopIteration         = Togo(C.PyExc_StopIteration)
	PyExc_SyntaxError           = Togo(C.PyExc_SyntaxError)
	PyExc_SystemError           = Togo(C.PyExc_SystemError)
	PyExc_SystemExit            = Togo(C.PyExc_SystemExit)
	PyExc_TabError              = Togo(C.PyExc_TabError)
	PyExc_TimeoutError          = Togo(C.PyExc_TimeoutError)
	PyExc_TypeError             = Togo(C.PyExc_TypeError)
	PyExc_UnboundLocalError     = Togo(C.PyExc_UnboundLocalError)
	PyExc_UnicodeDecodeError    = Togo(C.PyExc_UnicodeDecodeError)
	PyExc_UnicodeEncodeError    = Togo(C.PyExc_UnicodeEncodeError)
	PyExc_UnicodeError          = Togo(C.PyExc_UnicodeError)
	PyExc_UnicodeTranslateError = Togo(C.PyExc_UnicodeTranslateError)
	PyExc_ValueError            = Togo(C.PyExc_ValueError)
	PyExc_ZeroDivisionError     = Togo(C.PyExc_ZeroDivisionError)
)

//PyErr_NewException : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_NewException
func PyErr_NewException(name string, base, dict *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Togo(C.PyErr_NewException(cname, Toc(base), Toc(dict)))
}

//PyErr_NewExceptionWithDoc : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_NewExceptionWithDoc
func PyErr_NewExceptionWithDoc(name, doc string, base, dict *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cdoc := C.CString(doc)
	defer C.free(unsafe.Pointer(cdoc))

	return Togo(C.PyErr_NewExceptionWithDoc(cname, cdoc, Toc(base), Toc(dict)))
}

//PyException_GetTraceback : https://docs.python.org/3/c-api/exceptions.html#c.PyException_GetTraceback
func PyException_GetTraceback(ex *PyObject) *PyObject {
	return Togo(C.PyException_GetTraceback(Toc(ex)))
}

//PyException_SetTraceback : https://docs.python.org/3/c-api/exceptions.html#c.PyException_SetTraceback
func PyException_SetTraceback(ex, tb *PyObject) int {
	return int(C.PyException_SetTraceback(Toc(ex), Toc(tb)))
}

//PyException_GetContext : https://docs.python.org/3/c-api/exceptions.html#c.PyException_GetContext
func PyException_GetContext(ex *PyObject) *PyObject {
	return Togo(C.PyException_GetContext(Toc(ex)))
}

//PyException_SetContext : https://docs.python.org/3/c-api/exceptions.html#c.PyException_SetContext
func PyException_SetContext(ex, ctx *PyObject) {
	C.PyException_SetContext(Toc(ex), Toc(ctx))
}

//PyException_GetCause : https://docs.python.org/3/c-api/exceptions.html#c.PyException_GetCause
func PyException_GetCause(ex *PyObject) *PyObject {
	return Togo(C.PyException_GetCause(Toc(ex)))
}

//PyException_SetCause : https://docs.python.org/3/c-api/exceptions.html#c.PyException_SetCause
func PyException_SetCause(ex, cause *PyObject) {
	C.PyException_SetCause(Toc(ex), Toc(cause))
}
