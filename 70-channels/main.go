package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg *sync.WaitGroup = new(sync.WaitGroup)
)

func main() {
	var ch chan int // nil
	if ch == nil {
		println("ch is nil")
	}
	ch = make(chan int) // instantiate a unbuffered channel
	wg.Add(1)
	go func() {
		//time.Sleep(2 * time.Second)
		v := <-ch // receiving data from the chaannel
		time.Sleep(time.Second * 1)
		println(v)
		wg.Done()
	}()
	//time.Sleep(5 * time.Second)
	ch <- 100 ////-----         // sending data to chan
	fmt.Println("Done.. the receiver received")
	wg.Wait()
}

// func Fatal(v ...any) {
// fmt.Println(v...)
// os.Exit(1)
// }
//
// 1. Channel is used for communication between goroutines
// 2. Buffered and Unbuffered
// 3. use make built funciton to instantiate a channel
// 4.
