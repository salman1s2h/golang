package main

import (
	"fmt"
	"time"
)

func main() {
	ch, sig := make(chan int), make(chan struct{})
	go Sender(ch, 100)
	go Receiver(ch, sig, 100)

	<-sig
}

func Sender(ch chan<- int, n uint) {
	for i := 0; i <= int(n); i++ {
		time.Sleep(time.Millisecond * 50)
		ch <- i
	}
}

func Receiver(ch <-chan int, sig chan<- struct{}, n uint) {
	for i := 0; i <= int(n); i++ {
		fmt.Println(<-ch)
	}
	sig <- struct{}{}
}
