
#ifndef MACRO_H
#define MACRO_H

#include "Python.h"

int _go_Py_EnterRecursiveCall(const char *where);
void _go_Py_LeaveRecursiveCall();

int _go_PyType_Check(PyObject *o);

int _go_PyBool_Check(PyObject *o);
int _go_PyByteArray_Check(PyObject *o);

#endif