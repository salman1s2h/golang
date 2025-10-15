package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1, _ := Generator(2, 10, "gen-1", time.Millisecond*10)

	workers := uint8(20)
	wg := new(sync.WaitGroup)
	wg.Add(int(workers))
	go func() {
		for i := 0; i < int(workers); i++ {
			go func() {
				for c := range ch1 {
					//time.Sleep(time.Second * 1)
					println("worker-->"+fmt.Sprint(i), " ---->", c)
				}
				wg.Done()
			}()
		}
	}()
	//<-sig
	wg.Wait()

}

func Generator(mul uint, count uint, genname string, d time.Duration) (chan string, chan struct{}) {
	ch := make(chan string, 3)
	sig := make(chan struct{})
	go func() {
		for i := 1; i <= int(count); i++ {
			ch <- fmt.Sprint(genname, " sending data --", i*int(mul))
			time.Sleep(d)
		}
		close(ch)
		sig <- struct{}{}
	}()
	return ch, sig
}
