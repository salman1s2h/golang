package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1, sig1 := Generator(2, 10, "gen1", time.Millisecond*10)
	ch2, sig2 := Generator(3, 15, "gen2", time.Millisecond*20)
	ch3, sig3 := Generator(4, 20, "gen3", time.Millisecond*5)
	ch4, sig4 := Generator(5, 20, "gen4", time.Millisecond*100)

	wg := new(sync.WaitGroup)

	wg.Add(4)
	sig := make(chan struct{})
	// defer func() {
	// if r := recover(); r != nil {
	// fmt.Println(r)
	// }
	// }()
	go func() {
		wg.Wait()
		sig <- struct{}{}
	}()
out:
	for {
		select {
		case v, ok := <-ch1:
			if ok {
				println(v)
			}
		case v, ok := <-ch2:
			if ok {
				println(v)
			}
		case v, ok := <-ch3:
			if ok {
				println(v)
			}

		case v, ok := <-ch4:
			if ok {
				println(v)
			}

		case _, ok := <-sig1:
			if ok {
				wg.Done()
			}
		case _, ok := <-sig2:
			if ok {
				wg.Done()
			}
		case _, ok := <-sig3:
			if ok {
				wg.Done()
			}
		case _, ok := <-sig4:
			if ok {
				wg.Done()
			}
		case <-sig:
			break out
		}
	}
	//}()
}

func Generator(mul uint, count uint, genname string, d time.Duration) (chan string, chan struct{}) {
	ch := make(chan string)
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
