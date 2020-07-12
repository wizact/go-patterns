package main

import "fmt"

func main() {
	fmt.Println("Main() started")

	// Create a new channel of type string
	c := make(chan int, 3)

	go foo(c)

	c <- 1
	c <- 2
	c <- 3
	// Channel size is 3, so not until the 4th value the channel will overflow
	// If the below line is commented out, program exists without printing anything else
	// As channel overflows, it will be read until it is empty
	c <- 4

	// We can use cap and len functions to find the capacity and length of the channel
	fmt.Println(cap(c), len(c))

	fmt.Println("Main() stopped")
}

func foo(c chan int) {
	defer close(c)

	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num)
	}
}
