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

}

func TestReflectionLocals(t *testing.T) {

	locals := PyEval_GetLocals()
	assert.Nil(t, locals)

}

func TestReflectionGlobals(t *testing.T) {

	globals := PyEval_GetGlobals()
	assert.Nil(t, globals)

}

func TestReflectionFuncName(t *testing.T) {

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(len))

	assert.Equal(t, "len", PyEval_GetFuncName(len))

}
func TestReflectionFuncDesc(t *testing.T) {

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(len))

	assert.Equal(t, "()", PyEval_GetFuncDesc(len))
	Py_Finalize()
}
