package main

import (
	"errors"
	"fmt"
	"sync"
)

type CountType struct {
	C  int
	mu *sync.Mutex
	Wg *sync.WaitGroup
}

func NewCountType(c int, mu *sync.Mutex, wg *sync.WaitGroup) (*CountType, error) {
	if mu == nil {
		return nil, errors.New("nil mutex mu")
	}
	if wg == nil {
		return nil, errors.New("nil WaitGrouo Wg")
	}
	return &CountType{C: c, mu: mu, Wg: wg}, nil
}

func (c *CountType) Incr() {
	for i := 1; i <= 1000; i++ {
		c.mu.Lock()
		c.C++
		c.mu.Unlock()
	}
	c.Wg.Done()
}

func (c *CountType) Decr() {
	defer c.Wg.Done()
	for i := 1; i <= 1000; i++ {
		c.mu.Lock()
		c.C--
		c.mu.Unlock()
	}
}

func main() {

	ct, err := NewCountType(1, new(sync.Mutex), new(sync.WaitGroup))
	if err != nil {
		panic(err)
	}

	ct.Wg.Add(2)
	go ct.Incr()
	go ct.Decr()
	ct.Wg.Wait()
	fmt.Println(ct.C)
}
