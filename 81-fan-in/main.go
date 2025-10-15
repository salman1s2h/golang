package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := Generator(2, 10, "gen1", time.Millisecond*10)
	ch2 := Generator(3, 15, "gen2", time.Millisecond*20)
	ch3 := Generator(4, 20, "gen3", time.Millisecond*5)
	ch4 := Generator(5, 20, "gen4", time.Millisecond*100)
	out, sig := FanIn(ch1, ch2, ch3, ch4)

	go func() {
		for s := range out {
			fmt.Println(s)
		}
	}()

	<-sig

}

func Generator(mul uint, count uint, genname string, d time.Duration) chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= int(count); i++ {
			ch <- fmt.Sprint(genname, " sending data --", i*int(mul))
			time.Sleep(d)
		}
		close(ch)
	}()
	return ch
}

func FanIn(chs ...chan string) (chan string, chan struct{}) {
	out, sig := make(chan string), make(chan struct{})
	wg := new(sync.WaitGroup)

	merge := func(wg *sync.WaitGroup, ch chan string) {
		for c := range ch {
			out <- c
		}
		wg.Done()
	}

	for _, ch := range chs {
		wg.Add(1)
		// go func() {
		// for c := range ch {
		// out <- c
		// }
		// wg.Done()
		// }()
		go merge(wg, ch)
	}

	go func() {
		wg.Wait()
		sig <- struct{}{}
		close(out)
		close(sig)
	}()

	return out, sig
}
