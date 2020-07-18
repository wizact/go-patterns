package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(wg *sync.WaitGroup, instance int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done()
}

func main() {
	fmt.Println("main() started")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go foo(&wg, i)
	}

	wg.Wait()
	fmt.Println("main() stopped")
}
