package py35

// #include "Python.h"
import "C"

// Change : Removed 3.8 Onwards
//PyEval_ReInitThreads : https://docs.python.org/3/c-api/init.html#c.PyEval_ReInitThreads
func PyEval_ReInitThreads() {
	C.PyEval_ReInitThreads()
}

// Change : Removed 3.8 Onwards
//PyEval_InitThreads : https://docs.python.org/3/c-api/init.html#c.PyEval_InitThreads
func PyEval_InitThreads() {
	C.PyEval_InitThreads()
}

// Change : Removed 3.8 Onwards
//PyEval_ThreadsInitialized : https://docs.python.org/3/c-api/init.html#c.PyEval_ThreadsInitialized
func PyEval_ThreadsInitialized() bool {
	return C.PyEval_ThreadsInitialized() != 0
}
