package main

import (
	"time"
)

func main() {
	ch, sig := make(chan int), make(chan struct{})
	go Sender(ch, 10)
	go Receiver(ch, sig)

	<-sig
}

func Sender(ch chan<- int, n uint) {
	for i := 0; i <= int(n); i++ {
		time.Sleep(time.Millisecond * 50)
		ch <- i
	}
	close(ch) // close the channel is not make the channel nil
}

func Receiver(ch <-chan int, sig chan<- struct{}) {
	for v := range ch { // range loop tries to received the value until the channel is closed
		println(v)
	}
	sig <- struct{}{}
}
