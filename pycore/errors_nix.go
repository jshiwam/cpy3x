package pycore

/*
#include "Python.h"
*/
import "C"

//PySignal_SetWakeupFd : https://docs.python.org/3/c-api/exceptions.html#c.PySignal_SetWakeupFd
func PySignal_SetWakeupFd(fd uintptr) uintptr {
	return uintptr(C.PySignal_SetWakeupFd(C.int(fd)))
}
