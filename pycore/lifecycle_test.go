package pycore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialization(t *testing.T) {

	Py_Initialize()
	assert.True(t, Py_IsInitialized())
	Py_Finalize()
	assert.False(t, Py_IsInitialized())

}

// Not available in 3.5
// func TestInitializationEx(t *testing.T) {

// 	Py_Initialize()
// 	assert.True(t, Py_IsInitialized())
// 	// Change (Add) : version 3.6 onwards
// 	// assert.Zero(t, Py_FinalizeEx()
// 	assert.False(t, Py_IsInitialized())

// }

func TestProgramName(t *testing.T) {
	set := Py_SetStandardStreamEncoding("utf-8", "surrogateescape")
	assert.Zero(t, set)
	Py_Initialize()
	// Should call Py_GetProgramName after Py_Initialize
	defaultName, err := Py_GetProgramName()
	defer Py_SetProgramName(defaultName)

	assert.Nil(t, err)
	name := "py3é"
	Py_SetProgramName(name)
	newName, err := Py_GetProgramName()
	assert.Nil(t, err)
	assert.Equal(t, name, newName)

}

func TestPrefix(t *testing.T) {
	// Should call Py_GetPrefix after Py_Initialize
	prefix, err := Py_GetPrefix()
	assert.Nil(t, err)
	assert.IsType(t, "", prefix)

}

func TestExecPrefix(t *testing.T) {
	// Should call Py_GetExecPrefix after Py_Initialize
	execPrefix, err := Py_GetExecPrefix()
	assert.Nil(t, err)
	assert.IsType(t, "", execPrefix)

}

func TestProgramFullPath(t *testing.T) {
	// Should call Py_GetProgramFullPath after Py_Initialize
	programFullPath, err := Py_GetProgramFullPath()
	assert.Nil(t, err)
	assert.IsType(t, "", programFullPath)
}

func TestPythonHome(t *testing.T) {
	name := "høme"
	defaultHome, err := Py_GetPythonHome()
	defer Py_SetPythonHome(defaultHome)

	assert.Nil(t, err)
	Py_SetPythonHome(name)
	newName, err := Py_GetPythonHome()
	assert.Nil(t, err)
	assert.Equal(t, name, newName)
}

func TestPath(t *testing.T) {
	// Should call Py_GetPath after Py_Initialize
	defaultPath, err := Py_GetPath()
	assert.Nil(t, err)
	assert.IsType(t, "", defaultPath)
}

func TestVersion(t *testing.T) {
	version := Py_GetVersion()
	assert.IsType(t, "", version)
}

func TestPlatform(t *testing.T) {
	platform := Py_GetPlatform()
	assert.IsType(t, "", platform)
}

func TestCopyright(t *testing.T) {
	copyright := Py_GetCopyright()
	assert.IsType(t, "", copyright)
}

func TestCompiler(t *testing.T) {
	compiler := Py_GetCompiler()
	assert.IsType(t, "", compiler)
}

func TestBuildInfo(t *testing.T) {
	buildInfo := Py_GetBuildInfo()
	assert.IsType(t, "", buildInfo)
}

func TestSetArgv(t *testing.T) {

	PySys_SetArgv([]string{"test.py"})

	argv := PySys_GetObject("argv")
	assert.Equal(t, 1, PyList_Size(argv))
	assert.Equal(t, "test.py", PyUnicode_AsUTF8(PyList_GetItem(argv, 0)))

}

func TestSetArgvEx(t *testing.T) {

	PySys_SetArgvEx([]string{"test.py"}, false)

	argv := PySys_GetObject("argv")
	assert.Equal(t, 1, PyList_Size(argv))
	assert.Equal(t, "test.py", PyUnicode_AsUTF8(PyList_GetItem(argv, 0)))
	atexit := PyImport_ImportModule("atexit")
	atexit.DecRef()
	dthread := PyImport_AddModule("dummy_threading")
	defer dthread.DecRef()
	if dthread == nil {
		PyErr_Clear()
	}
	Py_Finalize()
}
