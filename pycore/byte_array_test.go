package pycore

import "testing"
import "github.com/stretchr/testify/assert"

func TestPyByteArrayCheck(t *testing.T) {
	Py_Initialize()

	s1 := "hWlloEorld"

	bytes := PyByteArray_FromStringAndSize(s1)
	defer DecRef(bytes)

	assert.True(t, PyByteArray_Check(bytes))
	assert.True(t, PyByteArray_CheckExact(bytes))

	Py_Finalize()
}

func TestPyByteArrayAsString(t *testing.T) {
	Py_Initialize()

	s2 := "hWlloEorld"

	bytes := PyByteArray_FromStringAndSize(s2)
	defer DecRef(bytes)

	assert.Equal(t, PyByteArray_AsString(bytes), s2)
	Py_Finalize()
}

func TestPyByteArrayResize(t *testing.T) {
	Py_Initialize()

	s3 := "hWlloEorld"
	newSize := 4
	bytes := PyByteArray_FromStringAndSize(s3)

	defer DecRef(bytes)

	assert.Equal(t, PyByteArray_Size(bytes), len(s3))
	// Resize to smaller size
	PyByteArray_Resize(bytes, newSize)

	assert.Equal(t, PyByteArray_Size(bytes), newSize)

	newSize = 20
	// Resize to bigger size
	PyByteArray_Resize(bytes, newSize)

	assert.Equal(t, PyByteArray_Size(bytes), newSize)

	Py_Finalize()
}

func TestPyByteArrayConcat(t *testing.T) {
	Py_Initialize()

	s4 := "eOrldHwllo"
	s5 := "hWlloEorld"

	a := PyByteArray_FromStringAndSize(s4)
	defer DecRef(a)

	b := PyByteArray_FromStringAndSize(s5)
	defer DecRef(b)

	ab := PyByteArray_Concat(a, b)
	defer DecRef(ab)

	result := PyByteArray_AsString(ab)

	assert.Equal(t, result, s4+s5)
	Py_Finalize()
}

func TestPyByteArrayFromObject(t *testing.T) {
	Py_Initialize()

	s6 := "WoRldoLLeH"

	bytes := PyBytes_FromString(s6)
	defer DecRef(bytes)

	assert.NotNil(t, bytes)

	bytearr := PyByteArray_FromObject(bytes)
	defer DecRef(bytearr)

	assert.NotNil(t, bytearr)

	Py_Finalize()
}
