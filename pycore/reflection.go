package pycore

/*
#include "Python.h"
*/
import "C"

//PyEval_GetBuiltins : https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetBuiltins
func PyEval_GetBuiltins() *PyObject {
	return Togo(C.PyEval_GetBuiltins())
}
