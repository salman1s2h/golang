package main

import (
	"fmt"
	"time"
)

func greet() {
	fmt.Println("Hello minds!")
}

func main() {

	go func() {
		fmt.Println("Hello World")
	}()

	go greet()

	fmt.Println("Main function")
	time.Sleep(time.Millisecond * 10)

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
