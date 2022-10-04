package py35

import (
	"github.com/jshiwam/cpy3x/pycore"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestClearFreeList(t *testing.T) {
	pycore.Py_Initialize()
	list := pycore.PyList_New(0)
	assert.True(t, pycore.PyList_Check(list))
	assert.True(t, pycore.PyList_CheckExact(list))
	defer list.DecRef()

	s := pycore.PyUnicode_FromString("world")
	assert.NotNil(t, s)

	list2 := pycore.PyList_New(1)
	defer list2.DecRef()

	assert.Zero(t, pycore.PyList_SetItem(list2, 0, s))

	assert.Zero(t, pycore.PyList_SetSlice(list, 0, 1, list2))

	list3 := pycore.PyList_GetSlice(list, 0, 1)
	assert.NotNil(t, list3)
	defer list3.DecRef()

	assert.Equal(t, 1, list2.RichCompareBool(list3, pycore.Py_EQ))

	tuple := pycore.PyList_AsTuple(list)
	assert.NotNil(t, tuple)
	defer tuple.DecRef()

	world := pycore.PyTuple_GetItem(tuple, 0)
	assert.NotNil(t, world)

	assert.Equal(t, "world", pycore.PyUnicode_AsUTF8(world))
	// Change : Removed 3.9 onwards
	PyList_ClearFreeList()
}

func TestClearFloatFreeList(t *testing.T) {
	pycore.Py_Initialize()

	assert.Equal(t, math.MaxFloat64, pycore.PyFloat_GetMax())

	assert.Equal(t, 2.2250738585072014e-308, pycore.PyFloat_GetMin())
	// Change : Removed 3.9 onwards
	PyFloat_ClearFreeList()
}

func TestClearDictFreeList(t *testing.T) {
	pycore.Py_Initialize()

	dict := pycore.PyDict_New()
	defer dict.DecRef()

	key1 := "key1"
	value1 := pycore.PyUnicode_FromString("value1")
	assert.NotNil(t, value1)
	defer value1.DecRef()

	key2 := pycore.PyUnicode_FromString("key2")
	assert.NotNil(t, key2)
	defer key2.DecRef()

	value2 := pycore.PyUnicode_FromString("value2")
	assert.NotNil(t, value2)
	defer value2.DecRef()

	err := pycore.PyDict_SetItem(dict, key2, value2)
	assert.Zero(t, err)

	err = pycore.PyDict_SetItemString(dict, key1, value1)
	assert.Zero(t, err)

	PyDict_ClearFreeList()

}
