package main

import (
	"fmt"
	"runtime"
	"time"
)

func greet() {
	fmt.Println("Hello minds!")
}

func main() {

	go func() {
		fmt.Println("Hello World")
		runtime.Goexit()
		fmt.Println("Hello World end")
	}()

	go greet()

	go func() {
		for i := 1; i <= 100; i++ {
			print(" ", i)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	fmt.Println("Main function")
	runtime.Goexit()
	//os.Exit(2)
	//os.Exit(2)

} // os.Exit(0)

// 1. go routine is a small thread/ green thread
// 2. main is also a go routine
// 3. no goroutine waits
// 		`for other goroutine to complete its execution
// 4. each goroutine size is ~2kb to begin with..
// 5. can seamlessly run few thousand go routines..
// 6. any function/method can be run as a goroutine
// 7. order of execution of goroutines cannot be guaranteed
// 8. There is no concept of parent or child go routines
