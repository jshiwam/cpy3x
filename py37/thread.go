package py37

// #cgo pkg-config: python37
// #include "Python.h"
import "C"

func PyEval_ReInitThreads() {
	C.PyEval_ReInitThreads()
}

//PyGILState_Check : https://docs.python.org/3/c-api/init.html#c.PyGILState_Check
func PyGILState_Check() bool {
	return C.PyGILState_Check() == 1
}

// Change : Removed 3.8 Onwards
//PyEval_InitThreads : https://docs.python.org/3/c-api/init.html#c.PyEval_InitThreads
// func PyEval_InitThreads() {
// 	C.PyEval_InitThreads()
// }

// Change : Removed 3.8 Onwards
//PyEval_ThreadsInitialized : https://docs.python.org/3/c-api/init.html#c.PyEval_ThreadsInitialized
// func PyEval_ThreadsInitialized() bool {
// 	return C.PyEval_ThreadsInitialized() != 0
// }
