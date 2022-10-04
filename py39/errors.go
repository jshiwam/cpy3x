//Change: Part Of Stable ABI since 3.6
//PyErr_SetImportErrorSubclass : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetImportErrorSubclass
// func PyErr_SetImportErrorSubclass(msg, name, path, subclass *PyObject) *PyObject {
// 	return Togo(C.PyErr_SetImportErrorSubclass(Toc(msg), Toc(name), Toc(path), Toc(subclass)))
// }