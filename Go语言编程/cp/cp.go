package main

/**
#include <stdio.h>
*/


import (
	"fmt"
	"C"
	"unsafe"
)

func main() {
	cstr := C.CString("Hello World!")

	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}


