package pycore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflectionBuiltins(t *testing.T) {
	Py_Initialize()

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(len))
	Py_Finalize()
}

func TestReflectionLocals(t *testing.T) {
	Py_Initialize()

	locals := PyEval_GetLocals()
	assert.Nil(t, locals)
	Py_Finalize()
}

func TestReflectionGlobals(t *testing.T) {
	Py_Initialize()

	globals := PyEval_GetGlobals()
	assert.Nil(t, globals)
	Py_Finalize()
}

func TestReflectionFuncName(t *testing.T) {
	Py_Initialize()

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(len))

	assert.Equal(t, "len", PyEval_GetFuncName(len))
	Py_Finalize()
}
func TestReflectionFuncDesc(t *testing.T) {
	Py_Initialize()

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(len))

	assert.Equal(t, "()", PyEval_GetFuncDesc(len))
	Py_Finalize()
}
