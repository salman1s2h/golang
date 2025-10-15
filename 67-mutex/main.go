package main

import (
	"fmt"
	"sync"
)

// Do not communicate by sharing memory; instead, share memory by communicating
var (
	Counter int             = 0
	mu      *sync.Mutex     = new(sync.Mutex)
	wg      *sync.WaitGroup = new(sync.WaitGroup)
)

type CountType struct {
	C  int
	mu *sync.Mutex
}

func (c *CountType) Incr() {
	for i := 1; i <= 1000; i++ {
		c.mu.Lock()
		Counter++
		c.mu.Unlock()
	}
}

func (c *CountType) Decr() {
	for i := 1; i <= 1000; i++ {
		c.mu.Lock()
		Counter--
		c.mu.Unlock()
	}
}

func Decr() {
	for i := 1; i <= 1000; i++ {
		mu.Lock()
		Counter--
		mu.Unlock()
	}
}

func main() {
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		Incr()
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		Decr()
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(Counter)
}

func Incr() {
	for i := 1; i <= 1000; i++ {
		mu.Lock()
		Counter++
		mu.Unlock()
	}
}
