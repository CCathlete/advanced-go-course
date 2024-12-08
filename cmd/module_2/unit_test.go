package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("expected %d but got %d", expected, result)
	}
}
