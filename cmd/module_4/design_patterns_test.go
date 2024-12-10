package main

import (
	"fmt"
	"testing"
)

/*
It's possible to have a type private and get an instance of it
using a function. It;s inner field SomeField can be public and modified
directly or it can be private and modified by a public function I
define.
*/
func TestGetInstance(t *testing.T) {
	t.Run("First time getting an instance", func(t *testing.T) {
		instance := GetInstance()
		instance.SomeField = "Hemlo"
		fmt.Println(instance.SomeField)
	})

	t.Run("Second time getting an instance", func(t *testing.T) {
		instance := GetInstance()
		instance.SomeField = "Hemlo"
		fmt.Println(instance.SomeField)
	})
}
