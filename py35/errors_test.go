
// Change: TypeError in 3.6 and 3.7
// func TestErrorInterrupt(t *testing.T) {
// 	Py_Initialize()

// 	PyErr_SetInterrupt()

// 	assert.Equal(t, -1, PyErr_CheckSignals())

// 	exc := PyErr_Occurred()

// 	assert.True(t, PyErr_GivenExceptionMatches(exc, PyExc_KeyboardInterrupt))

// 	assert.NotNil(t, PyErr_Occurred())
// 	PyErr_Clear()
// 	assert.Nil(t, PyErr_Occurred())
// }
