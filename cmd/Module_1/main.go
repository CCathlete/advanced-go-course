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

func Ex1() {
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

func Ex2() {
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

func Ex3() {
	original := "hello"

	reflected := reflect.ValueOf(original)
	reflectedInterface := reflected.Interface()

	twiceReflected := reflect.ValueOf(reflectedInterface)

	condition := twiceReflected.Interface() == original
	fmt.Printf("Is the twice reflected same as one? %v", condition)
}

func Check1() {
	// a := interface{}("hello")
	var a interface{} = "hello"
	fmt.Println(a)
	fmt.Printf("tyoe of a: %T", a)
}

func Ex4() {
	person := Person{
		Name:  "Ken",
		Age:   37,
		Email: "ken@example.com",
	}
	valueOfPerson := reflect.ValueOf(person)
	nameField := valueOfPerson.FieldByName("Name")
	fmt.Printf("Name field: %v\n", nameField.Interface())
	ageField := valueOfPerson.FieldByName("Age")
	fmt.Printf("Unmodified age: %v\n", ageField.Interface())
	ageField.SetInt(30) // I think this gives an error because
	// we can set the underlying value only if the reflect.Value is a pointer.
	fmt.Printf("Modified age: %v\n", ageField.Interface())
}

func Ex5() {
	data := interface{}(42)
	dataType := reflect.TypeOf(data)

	if dataType.Kind() == reflect.Int {
		intValue := data.(int)
		fmt.Printf("Data value: %v\nData type: %T\n", data, data)
		fmt.Printf("Int value: %v\n", intValue)
	}
}

func Ex6() {
	// typeName := "int"
	typeToCreate := reflect.TypeOf(0)
	if typeToCreate.Kind() == reflect.Int {
		newInt := reflect.New(typeToCreate).Elem().Interface().(int)
		fmt.Printf("New int instance: %v\n", newInt)
	}
}

func Ex7() {
	tt := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(0.0),
			Tag:  `json:"height"`,
		},
	})

	fmt.Println(tt.FieldByIndex([]int{0}))
	fmt.Println(tt.Field(0))
}

func main() {
	Ex7()
	// Check1()
}
