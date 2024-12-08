package main

import "fmt"

func Ex3() {
	type MyError struct {
		Code    int
		Message string
	}

	Error := func(e MyError) string {
		return fmt.Sprintf("Error %d: %s\n", e.Code, e.Message)
	}

	fmt.Println(Error(MyError{
		Code:    404,
		Message: "Hemlo from error land.",
	}))
}
