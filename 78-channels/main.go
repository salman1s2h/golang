package main

import (
	"sync"
	"time"
)

var wg *sync.WaitGroup = new(sync.WaitGroup)

func main() {
	ch := make(chan int)
	wg.Add(3)
	go Sender(wg, ch, 10)
	go Sender(wg, ch, 10)
	go Sender2(wg, ch, 10)
	go func() {
		wg.Wait()
		close(ch)
	}()
	sig := Receiver(ch)
	<-sig
	sig1 := Receiver(ch)
	<-sig1
	//<-Receiver(ch)
}

func Sender(wg *sync.WaitGroup, ch chan<- int, n uint) {
	for i := 0; i <= int(n); i++ {
		time.Sleep(time.Millisecond * 50)
		ch <- i
	}
	wg.Done()
	//close(ch) // close the channel is not make the channel nil
}
func Sender2(wg *sync.WaitGroup, ch chan<- int, n uint) {
	for i := 0; i <= int(n); i++ {
		time.Sleep(time.Millisecond * 50)
		ch <- i * i
	}
	wg.Done()
	//close(ch) // close the channel is not make the channel nil
}

func Receiver(ch <-chan int) chan struct{} {
	sig := make(chan struct{})
	go func() {
		for v := range ch { // range loop tries to received the value until the channel is closed
			println(v)
		}
		sig <- struct{}{}
	}()
	return sig
}
