package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

/*
Profiling the booting of a local server.
To view all results go to: localhost:3030/debug/pprof
To view the results graph run go tool pprof (it's a cli)
You need to create a server like I did here in Ex6
and get the profile data archive by a get request(or going to the url
with the browser) to: localhost:3030/debug/pprof/profile
port defined by you.
You then cd to the location of the prof archive and
run "go tool pprof".

To dowlnoad heap profiling results archive go to
localhost:3030/debug/pprof/heap
*/
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
