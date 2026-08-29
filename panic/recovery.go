package main

import (
	"flag"
	"fmt"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v", r)
		}
	}()

	// parse command-line flags
	name := flag.String("name", "default", "Name of the user")

	flag.Parse()

	// if the name is panic, then panic
	if *name == "panic" {
		panic("panic triggered")
	}

}
