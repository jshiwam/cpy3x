
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

// Should be in version 3.6 onwards
// func TestErrorImportErrorSubclass(t *testing.T) {
// 	Py_Initialize()

// 	message := PyUnicode_FromString("test message")
// 	defer message.DecRef()

// 	PyErr_SetImportErrorSubclass(message, nil, nil, Dict)

// 	assert.NotNil(t, PyErr_Occurred())
// 	PyErr_Clear()
// 	assert.Nil(t, PyErr_Occurred())
// }