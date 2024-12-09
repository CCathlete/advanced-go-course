package main

import (
	"log"
	"sync"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	wg := new(sync.WaitGroup) // returns initialised *sync.WaitGroup
	wg.Add(3)
	db, err := ConnectToDB()
	if err != nil {
		log.Fatalln(err)
	}

	go Ex1(db, wg)
	go Ex2(db, wg)
	go Ex3(wg)

	wg.Wait()
}
