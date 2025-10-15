package main

import (
	"runtime"
	"time"
)

func main() {
	ch1 := SqGenerater(time.Millisecond * 100)
	<-Receiver(ch1)

	ch2 := SqGenerater(time.Millisecond * 100)
	<-Receiver(ch2)
}

func SqGenerater(d time.Duration) chan int {
	ch := make(chan int)
	//timeOut := TimeOut(d)
	timeOut := time.After(d)
	go func() {
		i := 1
		for {
			select {
			case <-timeOut:
				close(ch)
				runtime.Goexit()
			default:
				ch <- i * i
				i++
			}
		}
	}()
	return ch
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

func TimeOut(d time.Duration) chan struct{} {
	sig := make(chan struct{})
	go func() {
		time.Sleep(d)
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}
