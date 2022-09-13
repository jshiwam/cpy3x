package pycore

// #include "Python.h"
import "C"

// Following APIs are the part of limited C API.

// PyThreadState
type PyThreadState C.PyThreadState

// PyGILState_STATE
type PyGILState C.PyGILState_STATE

// PyEval_InitThreads
func PyEval_InitThreads() {
	C.PyEval_InitThreads()
}

// PyEval_ThreadsInitialized()
func PyEval_ThreadsInitialized() bool {
	return C.PyEval_ThreadsInitialized() != 0
}

// PyEval_SaveThread()
func PyEval_SaveThread() *PyThreadState {
	return (*PyThreadState)(C.PyEval_SaveThread())
}

// PyEval_RestoreThread()
func PyEval_RestoreThread(tstate *PyThreadState) {
	C.PyEval_RestoreThread((*C.PyThreadState)(tstate))
}

// PyEval_ReleaseThread: https://docs.python.org/3.10/c-api/init.html#c.PyEval_ReleaseThread
func PyEval_ReleaseThread(tstate *PyThreadState) {
	C.PyEval_ReleaseThread((*C.PyThreadState)(tstate))
}

// PyEval_AcquireThread:
func PyEval_AcquireThread(tstate *PyThreadState) {
	C.PyEval_AcquireThread((*C.PyThreadState)(tstate))
}

// PyThreadState_Get()
func PyThreadState_Get() *PyThreadState {
	return (*PyThreadState)(C.PyThreadState_Get())
}

// PyThreadState_Swap
func PyThreadState_Swap(tstate *PyThreadState) *PyThreadState {
	return (*PyThreadState)(C.PyThreadState_Swap((*C.PyThreadState)(tstate)))
}

// PyGILState_Ensure
func PyGILState_Ensure() PyGILState {
	return PyGILState(C.PyGILState_Ensure())
}

// PyGILState_Release
func PyGILState_Release(state PyGILState) {
	C.PyGILState_Release(C.PyGILState_STATE(state))
}

// PyGILState_GetThisThreadState
func PyGILState_GetThisThreadState() *PyThreadState {
	return (*PyThreadState)(C.PyGILState_GetThisThreadState())
}
