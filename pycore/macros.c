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

// PyBool_Check : 
// Not a part of Stable ABI but present in versions 3.5-3.10
int _go_PyBool_Check(PyObject *o){
    return PyBool_Check(o);
}

// PyByteArray_Check : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Check
// Not a part of Stable ABI but present in versions 3.5-3.10
int _go_PyByteArray_Check(PyObject *o){
    return PyByteArray_Check(o);
}

// PyByteArray_CheckExact : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_CheckExact
// Not a part of Stable ABI but present in versions 3.5-3.10
int _go_PyByteArray_CheckExact(PyObject *o){
    return PyByteArray_CheckExact(o);
}

// PyObject_DelAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_DelAttr
// Not a part of Stable ABI but present in versions 3.5-3.10
int _go_PyObject_DelAttr(PyObject *o, PyObject *attr_name) {
    return PyObject_DelAttr(o, attr_name);
}

// PyObject_DelAttrString : https://docs.python.org/3/c-api/object.html#c.PyObject_DelAttrString
// Not a part of Stable ABI but present in versions 3.5-3.10
int _go_PyObject_DelAttrString(PyObject *o, const char *attr_name) {
    return PyObject_DelAttrString(o, attr_name);
}


int _go_PyDict_Check(PyObject *p) {
    return PyDict_Check(p);
}

int _go_PyDict_CheckExact(PyObject *p) {
    return PyDict_CheckExact(p);
}

int _go_PyType_CheckExact(PyObject *o) {
    return PyType_CheckExact(o);
}