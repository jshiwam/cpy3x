package pycore

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestThreadInitialization(t *testing.T) {
// 	// Change : Removed 3.9 Onwards
// 	Py_Initialize()
// 	// Change : Removed 3.9 Onwards
// 	PyEval_InitThreads()

// 	assert.True(t, PyEval_ThreadsInitialized())
// 	// Change : Removed 3.8 Onwards
// 	// PyEval_ReInitThreads()

// }

// func TestGIL(t *testing.T) {
// 	Py_Initialize()
// 	PyEval_InitThreads()

// 	gil := PyGILState_Ensure()

// 	assert.True(t, PyGILState_Check())

// 	PyGILState_Release(gil)
// }

// func TestThreadState(t *testing.T) {
// 	Py_Initialize()
// 	PyEval_InitThreads()

// 	threadState := PyGILState_GetThisThreadState()

// 	threadState2 := PyThreadState_Get()

// 	assert.Equal(t, threadState, threadState2)

// 	threadState3 := PyThreadState_Swap(threadState)

// 	assert.Equal(t, threadState, threadState3)
// }

// func TestThreadSaveRestore(t *testing.T) {
// 	Py_Initialize()
// 	PyEval_InitThreads()

// 	threadState := PyEval_SaveThread()

// 	assert.False(t, PyGILState_Check())

// 	PyEval_RestoreThread(threadState)
// }
