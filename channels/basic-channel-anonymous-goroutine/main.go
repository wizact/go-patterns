package main

import "fmt"

func main() {
	c := make(chan int)

	// just an anonymous goroutine in the wild
	go func(c chan int) {
		fmt.Println(<-c)
	}(c)

	c<- 1
}
