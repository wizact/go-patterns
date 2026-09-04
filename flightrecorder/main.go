package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	cfg := trace.FlightRecorderConfig{
		MinAge:   5 * time.Second,
		MaxBytes: 3 << 20, // 3 MB
	}
	fr := trace.NewFlightRecorder(cfg)
	fr.Start()
	defer fr.Stop()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range 10 {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Iteration %d\n", i)
		}
	}()
	<-done

	file, _ := os.Create("/tmp/trace.out")
	defer file.Close()
	fr.WriteTo(file)
}
