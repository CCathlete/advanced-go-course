package main

/*
	#include <stdio.h>
	void myCFunction() {
	printf("Hello from C function!\n");
	}
*/
import "C"

func Ex1() {
	C.myCFunction()
}
