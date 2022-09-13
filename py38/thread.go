package py38

// #cgo pkg-config: python3
// #include "Python.h"
import "C"

//PyGILState_Check : https://docs.python.org/3/c-api/init.html#c.PyGILState_Check
func PyGILState_Check() bool {
	return C.PyGILState_Check() == 1
}
