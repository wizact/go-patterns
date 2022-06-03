package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := errgroup.Group()
	g.Go(func() error {
		printAsync("Hello", time.Millisecond*100)
		return nil
	})
	g.Go(func() error {
		return fmt.Errorf("Errr!")
	})
	g.Go(func() error {
		printAsync("Bye", time.Millisecond*100)
		return nil
	})

	err := g.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error found %s", err.Error())
		os.Exit(1)
	}
	fmt.Print("All work completed as expected")
}

func printAsync(msg string, duration time.Duration) {
	time.Sleep(duration)
	fmt.Printf(msg)
}
