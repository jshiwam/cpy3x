// Change : Removed 3.9 onwards
//PyFloat_ClearFreeList : https://docs.python.org/3/c-api/float.html#c.PyFloat_ClearFreeList
// func PyFloat_ClearFreeList() int {
// 	return int(C.PyFloat_ClearFreeList())
// }

// https://docs.python.org/3/c-api/gcsupport.html?highlight=pygc_collect#c.PyGC_Collect