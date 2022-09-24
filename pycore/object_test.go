package pycore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttrString(t *testing.T) {
	Py_Initialize()

	sys := PyImport_ImportModule("sys")
	defer sys.DecRef()

	assert.True(t, sys.HasAttrString("stdout"))
	stdout := sys.GetAttrString("stdout")
	assert.NotNil(t, stdout)

	assert.Zero(t, sys.DelAttrString("stdout"))

	assert.Nil(t, sys.GetAttrString("stdout"))
	PyErr_Clear()

	assert.Zero(t, sys.SetAttrString("stdout", stdout))
}

func TestAttr(t *testing.T) {
	Py_Initialize()

	name := PyUnicode_FromString("stdout")
	defer name.DecRef()

	sys := PyImport_ImportModule("sys")
	defer sys.DecRef()

	assert.True(t, sys.HasAttr(name))
	stdout := sys.GetAttr(name)
	assert.NotNil(t, stdout)

	assert.Zero(t, sys.DelAttr(name))

	assert.Nil(t, sys.GetAttr(name))
	PyErr_Clear()

	assert.Zero(t, sys.SetAttr(name, stdout))
}

func TestRichCompareBool(t *testing.T) {
	Py_Initialize()

	s1 := PyUnicode_FromString("test1")
	s2 := PyUnicode_FromString("test2")

	assert.Zero(t, s1.RichCompareBool(s2, Py_EQ))

	assert.NotZero(t, s1.RichCompareBool(s1, Py_EQ))
}

func TestRichCompare(t *testing.T) {
	Py_Initialize()

	s1 := PyUnicode_FromString("test1")
	s2 := PyUnicode_FromString("test2")

	b1 := s1.RichCompare(s2, Py_EQ)
	defer b1.DecRef()
	assert.Equal(t, Py_False, b1)

	b2 := s1.RichCompare(s1, Py_EQ)
	assert.Equal(t, Py_True, b2)
	defer b2.DecRef()
}

func TestRepr(t *testing.T) {
	Py_Initialize()

	list := PyList_New(0)
	defer list.DecRef()

	repr := list.Repr()

	assert.Equal(t, "[]", PyUnicode_AsUTF8(repr))
}

func TestStr(t *testing.T) {
	Py_Initialize()

	list := PyList_New(0)
	defer list.DecRef()

	str := list.Str()

	assert.Equal(t, "[]", PyUnicode_AsUTF8(str))
}

func TestASCII(t *testing.T) {
	Py_Initialize()

	list := PyList_New(0)
	defer list.DecRef()

	ascii := list.ASCII()

	assert.Equal(t, "[]", PyUnicode_AsUTF8(ascii))
}

func TestCallable(t *testing.T) {
	builtins := PyEval_GetBuiltins()
	assert.True(t, PyDict_Check(builtins))
}
