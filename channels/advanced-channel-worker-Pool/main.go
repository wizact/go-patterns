package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond)
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num * num
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < 3; i++ {
		wg.Add(i)
		go worker(&wg, tasks, results, i)
	}

	for i := 0; i < 5; i++ {
		tasks <- i * 2
	}

	fmt.Println("[main] wrote 5 tasks")
	close(tasks)

	wg.Wait()

	for i := 0; i < 5; i++ {
		result := <-results
		fmt.Println("[main] result", i, ":", result)
	}
}
