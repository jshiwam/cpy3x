package py35

import (
	"github.com/jshiwam/cpy3x/pycore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreadInitialization(t *testing.T) {
	pycore.Py_Initialize()

	PyEval_InitThreads()

	assert.True(t, PyEval_ThreadsInitialized())
	PyEval_ReInitThreads()

}

func TestGIL(t *testing.T) {
	pycore.Py_Initialize()
	PyEval_InitThreads()

	gil := pycore.PyGILState_Ensure()

	assert.True(t, pycore.PyGILState_Check())

	pycore.PyGILState_Release(gil)
}

func TestThreadState(t *testing.T) {
	pycore.Py_Initialize()
	PyEval_InitThreads()

	threadState := pycore.PyGILState_GetThisThreadState()

	threadState2 := pycore.PyThreadState_Get()

	assert.Equal(t, threadState, threadState2)

	threadState3 := pycore.PyThreadState_Swap(threadState)

	assert.Equal(t, threadState, threadState3)
}

func TestThreadSaveRestore(t *testing.T) {
	pycore.Py_Initialize()
	PyEval_InitThreads()

	threadState := pycore.PyEval_SaveThread()

	assert.False(t, pycore.PyGILState_Check())

	pycore.PyEval_RestoreThread(threadState)
}
