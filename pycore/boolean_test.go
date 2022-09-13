package pycore

import "testing"
import "github.com/stretchr/testify/assert"

func TestBoolFromLong(t *testing.T) {
	Py_Initialize()
	assert.Equal(t, Py_True, PyBool_FromLong(1))
	assert.Equal(t, Py_False, PyBool_FromLong(0))
	Py_Finalize()
}
