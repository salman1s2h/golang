package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := Generator(2, 10, "gen1", time.Millisecond*10)
	ch2 := Generator(3, 15, "gen2", time.Millisecond*20)
	ch3 := Generator(4, 20, "gen3", time.Millisecond*5)
	ch4 := Generator(5, 20, "gen4", time.Millisecond*100)

	// wg := new(sync.WaitGroup)
	// wg.Add(4)

	//go func() {
	counter := uint8(0)
	for {
		if counter == 4 {
			break
		}
		select {
		case v, ok := <-ch1:
			if ok {
				println(v)
			} else {
				counter++
				//wg.Done()
			}

		case v, ok := <-ch2:
			if ok {
				println(v)
			} else {
				counter++
				//wg.Done()
			}
		case v, ok := <-ch3:
			if ok {
				println(v)
			} else {
				counter++
				//wg.Done()
			}
		case v, ok := <-ch4:
			if ok {
				println(v)
			} else {
				counter++
				//wg.Done()
			}
			// default:
			// println("lot of disturbance")
		}
	}
	//}()

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
