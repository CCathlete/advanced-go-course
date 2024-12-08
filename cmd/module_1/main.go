package main

import (
	"fmt"
	"reflect"
	"unsafe"
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

func Ex8() {
	var x int = 42
	y := *(*float64)(unsafe.Pointer(&x))
	fmt.Printf("x (int): %d\n", x)
	fmt.Printf("y (float64): %f\n", y)
}

func Ex9() {
	person := Person{
		Name: "Ken",
		Age:  37,
	}

	size := unsafe.Sizeof(person)
	addr := uintptr(unsafe.Pointer(&person))
	nameOffset := uintptr(unsafe.Offsetof(person.Name))
	ageOffset := uintptr(unsafe.Offsetof(person.Age))
	namePtr := (*string)(unsafe.Pointer(addr + nameOffset))
	agePtr := (*int)(unsafe.Pointer(addr + ageOffset))

	fmt.Printf("Memory layout of Person struct (size=%d):\n", size)
	fmt.Printf("Name: %s\n", *namePtr)
	fmt.Printf("Age: %d\n", *agePtr)
}

func Ex10() {
	var i interface{} = 42
	y := i.(int) // Asserted i turns it to a real int.
	fmt.Printf("Size of unasserted: %d\n", unsafe.Sizeof(i))
	fmt.Printf("Size of asserted: %d\n", unsafe.Sizeof(y))
	tyI := reflect.TypeOf(i)
	// This shows the type of the underlying data, but i is an interface = an array of 2 pointers.
	// One to type info and the other to the value of the data.
	fmt.Printf("Type of unasserted's underlying data: %d\n", tyI.Kind())

	pi := unsafe.Pointer(&i) // Pointer to an interface!
	// pi is a pointer so we can convert it only to a pointer.
	// We convert it to a pointer of an array with size 2 and an underlying type of unsafe.Pointer.
	// We then dereference it and take the second value of the array which is an unsafe.Pointer to an int.
	convPi := (*[2]unsafe.Pointer)(pi)[1]
	valueConvPi := *(*int)(convPi)
	py := unsafe.Pointer(&y) // Pointer to an int64 (size of the pointer is 8 bytes = 64 bits)
	vy := *(*int)(py)
	fmt.Printf("value: %d\n", valueConvPi)
	fmt.Printf("value: %d\n", y)
	fmt.Printf("value: %v\n", vy)
}

func main() {
	Ex10()
	// Check1()
}
