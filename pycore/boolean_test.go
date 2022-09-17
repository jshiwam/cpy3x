package pycore

import "testing"
import "github.com/stretchr/testify/assert"

func TestBoolFromLong(t *testing.T) {
	Py_Initialize()
	assert.Equal(t, Py_True, PyBool_FromLong(1))
	assert.Equal(t, Py_False, PyBool_FromLong(0))
	Py_Finalize()
}

func TestBoolCheck(t *testing.T) {
	Py_Initialize()
	assert.True(t, PyBool_Check(Py_True))
	assert.True(t, PyBool_Check(Py_False))
	Py_Finalize()
}
