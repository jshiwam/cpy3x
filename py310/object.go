
// // Change: Part of Stable ABI Should Go In Version 3.10
// //GetAIter : https://docs.python.org/3/c-api/object.html#c.PyObject_GetAIter
// func (pyObject *PyObject) GetAIter() *PyObject {
// 	return Togo(C.PyObject_GetAIter(Toc(pyObject)))
// }
