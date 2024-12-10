package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"runtime/trace"
	"sync"
	"time"
)

const (
	NumWorkers    = 4
	NumTasks      = 500
	MemoryIntense = 10000 // Size of memory intensive tasks (number of elements).
)

func ManageGC() {
	f, _ := os.Create("trace.out") // Writing to a trace file.
	trace.Start(f)
	defer trace.Stop()

	// Setting the target precentage for the garbage collector. Default is 100%.
	debug.SetGCPercent(100)

	// Creating tasks and results queues.
	taskQueue := make(chan int, NumTasks)
	resultkQueue := make(chan int, NumTasks)

	// Start workers. A workers pool is a waitgroup with num of workers.
	var wg sync.WaitGroup
	wg.Add(NumWorkers)
	for i := 0; i < NumWorkers; i++ {
		go worker(taskQueue, resultkQueue, &wg)
	}

	// Sending tasks to the queue.
	for i := 0; i < NumTasks; i++ {
		taskQueue <- i
	}
	close(taskQueue)

	// Closing the results queue after all goroutines are done.
	go func() {
		wg.Wait()
		close(resultkQueue)
	}()

	// Processing the results.
	for result := range resultkQueue {
		fmt.Println("Result: ", result)
	}
	fmt.Println("Done!")

}

// Worker function.
func worker(tq <-chan int, rq chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tq {
		result := performMemoryIntensiveTask(task)
		rq <- result
	}
}

func performMemoryIntensiveTask(task int) int {
	data := make([]int, MemoryIntense)
	for i := 0; i < MemoryIntense; i++ {
		data[i] = i + task // Addition of large numbers.
	}

	time.Sleep(10 * time.Millisecond)
	// Calculate the result.
	result := 0
	for _, value := range data {
		result += value
	}

	return result
}
