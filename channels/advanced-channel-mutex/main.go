package main

import (
	"fmt"
	"sync"
)

var i int

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	i = i + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	wg.Wait()
	fmt.Println("i is", i)
}
