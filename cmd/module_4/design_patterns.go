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
