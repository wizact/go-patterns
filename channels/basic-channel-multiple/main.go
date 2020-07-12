package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main() started")

	// Create two new channel of type string
	fooChan := make(chan int)
	barChan := make(chan int)

	go foo(fooChan)
	go bar(barChan)

	testNum := 3

	fooChan <- testNum
	barChan <- testNum
	// main  is blocked here until goroutines read from the channels
	fmt.Println("main unblocked between end of read to start of write in goroutines")
	// main is blocked again waiting until goroutines write to the channels
	fooVal, barVal := <-fooChan, <-barChan

	fmt.Println(fooVal, barVal, fooVal + barVal)

	fmt.Println("Main() stopped")
}

func foo(c chan int) {
	fmt.Println("[foo] started")
	num := <-c
	c <- num
	fmt.Println("[foo] ended")
}

func bar(c chan int) {
	fmt.Println("[bar] started")
	num := <-c
	c <- -num
	fmt.Println("[bar] ended")
}
