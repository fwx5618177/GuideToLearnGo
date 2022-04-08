package main

/*
#include <stdio.h>

static void SayHi(const char* s) {
	puts(s);
}
*/
import "C"

func main() {
	C.SayHi(C.CString("Hello!\n"))
}