package py37

import "testing"
import "github.com/jshiwam/cpy3x/pycore"
import "github.com/stretchr/testify/assert"

func TestThreadInitialization(t *testing.T) {
	pycore.Py_Initialize()

	assert.True(t, pycore.PyEval_ThreadsInitialized())

	PyEval_ReInitThreads()
	pycore.Py_Finalize()
}

func TestGIL(t *testing.T) {
	pycore.Py_Initialize()

	gil := pycore.PyGILState_Ensure()

	assert.True(t, PyGILState_Check())

	pycore.PyGILState_Release(gil)
	pycore.Py_Finalize()
}

func TestThreadState(t *testing.T) {
	pycore.Py_Initialize()

	threadState := pycore.PyGILState_GetThisThreadState()
	threadState2 := pycore.PyThreadState_Get()

	assert.Equal(t, threadState, threadState2)

	threadState3 := pycore.PyThreadState_Swap(threadState)

	assert.Equal(t, threadState, threadState3)
	pycore.Py_Finalize()
}

func TestThreadSaveRestore(t *testing.T) {
	pycore.Py_Initialize()
	threadState := pycore.PyEval_SaveThread()
	assert.False(t, PyGILState_Check())
	pycore.PyEval_RestoreThread(threadState)
	assert.True(t, PyGILState_Check())
	pycore.Py_Finalize()
}
