#include "macros.h"

// Part of Stable ABI Since python-3.9
// Py_EnterRecursiveCall: https://docs.python.org/3.10/c-api/exceptions.html#c.Py_EnterRecursiveCall
int _go_Py_EnterRecursiveCall(const char *where){
    return Py_EnterRecursiveCall(where);
}

// Part of Stable ABI Since python-3.9
// Py_LeaveRecursiveCall: https://docs.python.org/3.10/c-api/exceptions.html#c.Py_LeaveRecursiveCall
void _go_Py_LeaveRecursiveCall(){
    Py_LeaveRecursiveCall();
}

// PyType_Check : https://docs.python.org/3.10/c-api/type.html?highlight=pytype_check#c.PyType_Check
// Not a part of Stable ABI but present in versions 3.5-3.10
int _go_PyType_Check(PyObject *o){
    return PyType_Check(o);
}

// Not a part of Stable ABI but present in versions 3.5-3.10
int _go_PyBool_Check(PyObject *o){
    return PyBool_Check(o);
}

// PyByteArray_Check : 
// Not a part of Stable ABI but present in versions 3.5-3.10
int _go_PyByteArray_Check(PyObject *o){
    return PyByteArray_Check(o);
}
