package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		fmt.Println("Hello")
	})

	wg.Go(func() {
		fmt.Println("Bye")
	})

	wg.Wait()

	fmt.Println("All work completed as expected")
}
