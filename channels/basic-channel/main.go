package main

import "fmt"

func main() {
	fmt.Println("Main() started")

	// Create a new channel of type string
	c := make(chan string)

	// Close the channel at the end of execution of this function
	defer close(c)

	go foo(c)

	// Send Hello to string channel (c)
	c <- "bar"

	fmt.Println("Main() stopped")
}

func foo(c chan string) {
	// Receive string value from channel c
	fmt.Println("Hello " + <-c)
}
