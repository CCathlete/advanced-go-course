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

func GetInstance() *mySingleton {
	once.Do(func() {
		fmt.Println("Allocating a new address for the instance.")
		instance = &mySingleton{}
	})
	return instance
}
