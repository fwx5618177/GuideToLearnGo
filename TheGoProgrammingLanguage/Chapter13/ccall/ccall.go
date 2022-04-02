package ccall

import "C"

import (
	"io"
	"unsafe"
)

func main() {
	C.hi()
}
