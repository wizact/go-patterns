package main

import "fmt"

func main() {
	fmt.Println("Main() started")

	// Create a new channel of type string
	c := make(chan int)

	go foo(c)

	// blocks/unblocks while waiting for a value from the channel
	for {
		val, ok := <-c
		if !ok {
			fmt.Println(val, ok, " <- Channel ended")
			break
		} else {
			fmt.Println(val, ok, " <- channel seems to be working")
		}

	}
	fmt.Println("Main() stopped")
}

func foo(c chan int) {
	defer close(c)

	for i := 0; i < 9; i++ {
		c <- i
	}
}
