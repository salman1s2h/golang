package main

import "fmt"

func main() {
	defer println("end of main")
	func() { // func-1
		defer println("end of func-1")
		func() { // func-1.1
			defer println("end of func-1.1")
			func() { // func-1.1.2
				fmt.Println("Before panic")
			}()
			n := 0
			println(100 / n)
			func() { // func 1.1.3
				defer println("end of func-1.1.3")
				fmt.Println("After panic")
			}()
		}()
		func() { // func 1.2
			fmt.Println("Hello World")
		}()
	}()

}

// 1. error
// 2. panic
// 3. fatal
