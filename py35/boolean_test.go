package py35

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/jshiwam/cpy3x/pycore"

func TestBoolFromLong(t *testing.T) {
	pycore.Py_Initialize()
	assert.Equal(t, pycore.Py_True, PyBool_FromLong(1))
	assert.Equal(t, pycore.Py_False, PyBool_FromLong(0))
	pycore.Py_Finalize()
}
