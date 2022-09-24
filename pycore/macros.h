
#ifndef MACRO_H
#define MACRO_H

#include "Python.h"

int _go_Py_EnterRecursiveCall(const char *where);
void _go_Py_LeaveRecursiveCall();

int _go_PyType_Check(PyObject *o);

int _go_PyBool_Check(PyObject *o);

int _go_PyByteArray_Check(PyObject *o);
int _go_PyByteArray_CheckExact(PyObject *o);

int _go_PyObject_DelAttr(PyObject *o, PyObject *attr_name);
int _go_PyObject_DelAttrString(PyObject *o, const char *attr_name);

int _go_PyDict_Check(PyObject *p);
int _go_PyDict_CheckExact(PyObject *p);

int _go_PyType_CheckExact(PyObject *o);
#endif