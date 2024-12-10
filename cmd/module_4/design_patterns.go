package main

import (
	"fmt"
	"sync"
)

// Singleton pattern - using package level variable as a use flag.
type mySingleton struct {
	SomeField string
}

var once sync.Once
var instance *mySingleton

// The once variable will make sure that the memory allocation
// would happen only once. The rest of the times the Get function
// is called, the same instance (with the same address) is returned.
func GetInstance() *mySingleton {
	once.Do(func() {
		fmt.Println("Allocating a new address for the instance.")
		instance = &mySingleton{}
	})
	return instance
}

//-----------------------------------------------------------------------------------------------------------------------

// Polymorphism + composition.

type PlainPizza struct {
	Dough string
}

func (p *PlainPizza) Cost() int {
	return 10
}

type SpecialPizza struct {
	Dough      string
	SpecialMod []string // Special modifications to the dough.
}

func (p *SpecialPizza) Cost() int {
	return 20
}

type Pizza interface {
	Cost() int
}

type PizzaDecorator struct { // Plain + stuff
	Base     Pizza // Could be plain or special pizza.
	Sauce    string
	Toppings []string
}

type CheeseTopping struct {
	PizzaDecorator
	CheeseType string
}

func (c *CheeseTopping) Cost() int {
	return c.Base.Cost() + 2
}

// -----------------------------------------------------------------------------------------------------------------------

// Builder interface defines the steps to build the product.
// We also return the builder to be able to concatenate methods later
// (It's not critical since we pass a pointer to the pizza builder
// so it'll save the changes without returning).
type PizzaBuilder interface {
	SetDough() PizzaBuilder
	SetSauce() PizzaBuilder
	SetToppings() PizzaBuilder
	GetPizza() PizzaDecorator
}

// A concrete builder that implements the builder interface.
type MargheritaBuilder struct {
	pizza PizzaDecorator
}

// Returns an initialised instance.
func NewMargheritaBuilder() *MargheritaBuilder {
	return &MargheritaBuilder{pizza: PizzaDecorator{}}
}

func (b *MargheritaBuilder) SetDough() PizzaBuilder {
	var base PlainPizza
	base.Dough = "thin crust, wheat"
	b.pizza.Base = &base // Interface is like a pointer so we need &base.
	return b
}

func (b *MargheritaBuilder) SetSauce() PizzaBuilder {
	b.pizza.Sauce = "Tomato"
	return b
}

func (b *MargheritaBuilder) SetToppings() PizzaBuilder {
	b.pizza.Toppings = []string{"Mozzarella", "Basil"}
	return b
}

func (b *MargheritaBuilder) GetPizza() PizzaDecorator {
	return b.pizza
}

// -----------------------------------------------------------------------------------------------------------------------

// Director manages the building process.
type PizzaDirector struct {
	builder PizzaBuilder
}

func NewPizzaDirector(builder PizzaBuilder) *PizzaDirector {
	return &PizzaDirector{builder: builder}
}

func (d *PizzaDirector) ConstructPizza() {
	// d.builder = d.builder.SetDough().SetSauce().SetToppings()
	d.builder.SetDough().SetSauce().SetToppings()
}
