package main

/*
#cgo CFLAGS: -I./c_files
#cgo LDFLAGS: -L./c_files
#include "myFuncs.h"
#include <stdio.h>

void myCFunction() {
printf("Hello from C function!\n");
}
*/
import "C"

func Ex1() {
	C.myCFunction()
}

func Ex2() {
	C.someFunc()
}
