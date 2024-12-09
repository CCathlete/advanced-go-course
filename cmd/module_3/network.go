package main

import "sync"

func Ex3(wg *sync.WaitGroup) {
	defer wg.Done()
}
