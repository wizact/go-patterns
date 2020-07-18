package main

import "fmt"

func bar(c chan string) {
	fmt.Println(<-c)
}

func foo(cc chan chan string) {
	c := make(chan string)

	cc <- c
}

func main() {
	fmt.Println("main() started")

	cc := make(chan chan string)

	go foo(cc)

	c := <-cc

	go bar(c)

	c <- "baz"

	fmt.Println("main() stopped")
}
