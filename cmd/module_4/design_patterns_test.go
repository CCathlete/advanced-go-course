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

func TestGetMargherita(t *testing.T) {
	t.Run("Let's make a Margherita.", func(t *testing.T) {
		builder := NewMargheritaBuilder()
		director := NewPizzaDirector(builder)
		director.ConstructPizza()
		pizza := builder.pizza
		fmt.Printf(`Pizza details:
		Dough: %s
		Sauce: %s
		Toppings: %s
		Cost: %d
		`, pizza.Base.(*PlainPizza).Dough,
			pizza.Sauce,
			pizza.Toppings,
			pizza.Base.Cost())
	})
}

func TestGCManagement(t *testing.T) {
	t.Run("Testing management o GC with large storage and arithmetic.",
		func(t *testing.T) {
			ManageGC()
		})
}
