package main

import (
	"runtime"
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
	for {
		v, ok := <-ch
		if !ok {
			sig <- struct{}{}
			//return
			// break
			runtime.Goexit()
		} else {
			println(v)
		}
	}
}
