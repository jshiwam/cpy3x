package pycore

import "testing"
import "github.com/stretchr/testify/assert"

func TestPyByteArrayCheck(t *testing.T) {
	Py_Initialize()
	defer Py_Finalize()

	s1 := "hWlloEorld"

	bytes := PyByteArray_FromStringAndSize(s1)
	defer bytes.DecRef()

	assert.True(t, PyByteArray_Check(bytes))
	assert.True(t, PyByteArray_CheckExact(bytes))

}

func TestPyByteArrayAsString(t *testing.T) {
	Py_Initialize()
	defer Py_Finalize()
	s2 := "hWlloEorld"

	bytes := PyByteArray_FromStringAndSize(s2)
	defer bytes.DecRef()

	assert.Equal(t, PyByteArray_AsString(bytes), s2)

}

func TestPyByteArrayResize(t *testing.T) {
	Py_Initialize()
	defer Py_Finalize()

	s3 := "hWlloEorld"
	newSize := 4
	bytes := PyByteArray_FromStringAndSize(s3)

	defer bytes.DecRef()

	assert.Equal(t, PyByteArray_Size(bytes), len(s3))
	// Resize to smaller size
	PyByteArray_Resize(bytes, newSize)

	assert.Equal(t, PyByteArray_Size(bytes), newSize)

	newSize = 20
	// Resize to bigger size
	PyByteArray_Resize(bytes, newSize)

	assert.Equal(t, PyByteArray_Size(bytes), newSize)

}

func TestPyByteArrayConcat(t *testing.T) {
	Py_Initialize()
	defer Py_Finalize()

	s4 := "eOrldHwllo"
	s5 := "hWlloEorld"

	a := PyByteArray_FromStringAndSize(s4)
	defer a.DecRef()

	b := PyByteArray_FromStringAndSize(s5)
	defer b.DecRef()

	ab := PyByteArray_Concat(a, b)
	defer ab.DecRef()

	result := PyByteArray_AsString(ab)

	assert.Equal(t, result, s4+s5)

}

func TestPyByteArrayFromObject(t *testing.T) {
	Py_Initialize()
	defer Py_Finalize()

	s6 := "WoRldoLLeH"

	bytes := PyBytes_FromString(s6)
	defer bytes.DecRef()

	assert.NotNil(t, bytes)

	bytearr := PyByteArray_FromObject(bytes)
	defer bytearr.DecRef()

	assert.NotNil(t, bytearr)

}
