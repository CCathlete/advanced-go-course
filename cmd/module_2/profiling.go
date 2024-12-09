package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

// Profiling the booting of a local server.
// To view the results graph run go tool pprof (it's a cli)
func Ex6() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		log.Println("Booting on localhost:3030")
		log.Fatal(http.ListenAndServe(":3030", nil))
	}()
	wg.Wait()
}
