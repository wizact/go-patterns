package main

import "fmt"

func main() {
	fmt.Println("Main() started")

	// Create a new channel of type string
	c := make(chan int)

	go foo(c)

	// Using range c will automatically break when channel closes
	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("Main() stopped")
}

func foo(c chan int) {
	defer close(c)

	for i := 0; i < 9; i++ {
		c <- i
	}
}
