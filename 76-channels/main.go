package main

func main() {
	ch := FibGenerater()
	<-Receiver(ch)
}

func FibGenerater() chan int {
	ch := make(chan int)
	go func() {
		a, b := 0, 1
		for i := 0; i <= 10; i++ {
			ch <- a
			a, b = b, a+b
		}
		close(ch)
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
