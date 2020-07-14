package main

import "fmt"

func foo(c <-chan int) {
	fmt.Println(<-c)
}

func main() {
	// This channel is bi-directional but in foo func params we changed it to uni-directional
	c := make(chan int)

	go foo(c)

	c <- 1
}
