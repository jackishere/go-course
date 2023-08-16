package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Second) // Simulate some processing time
		results <- task * 2
	}
}
func main() {
	numWorkers := 4000000
	numTasks := 100

	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	var wg sync.WaitGroup

	// Create worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Enqueue tasks
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect and print results
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
	}
}
