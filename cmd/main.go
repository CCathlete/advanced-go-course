package main

import (
	"fmt"
	"reflect"
)

func observe(a any) {
	fmt.Printf("The type is: %T\n", a)
	fmt.Printf("The value is: %v\n", a)
	fmt.Println("------------------------------")
}

func ex1() {
	var value float64 = 25
	value2 := "welcome learners"
	observe(value)
	observe(value2)
	T := reflect.TypeOf(value)
	fmt.Println(T)
}

type Person struct {
	Name  string
	Age   int
	Email string
}

func ex2() {
	person := Person{
		Name:  "Ken",
		Age:   37,
		Email: "ken@example.com",
	}

	structType := reflect.TypeOf(person)
	fmt.Printf("Type of of struct: %s\n", structType)
	for i := 1; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fmt.Printf("Field name: %s. Field type: %s\n", field.Name, field.Type)
	}
}

func main() {
	ex2()
}
