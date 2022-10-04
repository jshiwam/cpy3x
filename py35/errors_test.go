package py35

import (
	"github.com/jshiwam/cpy3x/pycore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorInterrupt(t *testing.T) {
	pycore.Py_Initialize()

	pycore.PyErr_SetInterrupt()

	assert.Equal(t, -1, pycore.PyErr_CheckSignals())

	exc := pycore.PyErr_Occurred()

	assert.True(t, pycore.PyErr_GivenExceptionMatches(exc, pycore.PyExc_KeyboardInterrupt))

	assert.NotNil(t, pycore.PyErr_Occurred())
	pycore.PyErr_Clear()
	assert.Nil(t, pycore.PyErr_Occurred())
}
