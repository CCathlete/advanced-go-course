package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Table driven tests.
func TestEx5(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		fmt.Println("Addition.")
	})
	t.Run("Substraction", func(t *testing.T) {
		fmt.Println("Substraction.")
	})
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

// Test assertion.
func TestWithAssert(t *testing.T) {
	expected := "Hi I'm the expected.\n"
	result := func() string {
		return fmt.Sprintln("Just testing this with assertion.")
	}()
	assert.Equal(t, expected, result)
}
