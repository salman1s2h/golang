package main

import (
	"fmt"
	"sync"
)

var (
	wg *sync.WaitGroup = new(sync.WaitGroup) // instantiated it
)

func greet(wg *sync.WaitGroup) {
	fmt.Println("Hello minds!")
	wg.Done()
}

func main() {
	defer wg.Wait() // how long it waits, it waits until the counter/ state becomes zero
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello World")
		//runtime.Goexit()
		fmt.Println("Hello World end")
	}()

	wg.Add(1)
	go greet(wg)

	wg.Add(1)
	go func() {
		wg.Add(1)
		go func() {
			wg.Add(1)
			go func() {
				for i := 1; i <= 100; i++ {
					wg.Add(1)
					go func() {
						print(" ", i)
						wg.Done()
					}()
					//time.Sleep(time.Millisecond * 200)
				}
				wg.Done()
			}()
			wg.Done()
		}()
		wg.Done()

	}()
	fmt.Println("Main function")
}

//
//
//
//
//
//
//
//
//
