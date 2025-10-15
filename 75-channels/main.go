package main

import (
	"time"
)

func main() {
	ch := make(chan int)
	//go Sender(ch, 10)
	go Sender(ch, 10)

	sig := Receiver(ch)
	<-sig
	//<-Receiver(ch)
}

func Sender(ch chan<- int, n uint) {
	for i := 0; i <= int(n); i++ {
		time.Sleep(time.Millisecond * 50)
		ch <- i
	}
	close(ch) // close the channel is not make the channel nil
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
