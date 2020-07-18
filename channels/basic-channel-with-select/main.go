package main

import (
	"fmt"
	"time"
)

func foo(c chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		c <- "foo won"
	}
}

func bar(c chan string) {
	for i := 0; i < 2; i++ {
		time.Sleep(5 * time.Second)
		c <- "bar won"
	}
}

func main() {
	fmt.Println("main() started")

	chan1 := make(chan string)
	chan2 := make(chan string)

	go foo(chan1)
	go bar(chan2)

	for {
		select {
		case res := <-chan1:
			fmt.Println("foo says", res)
		case res := <-chan2:
			fmt.Println("bar says", res)
		default: // default makes all channels non-blocking
			break
		}
	}
}
