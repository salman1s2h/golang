package main

import "fmt"

func main() {
	// 1.22
	fmt.Println("print 1-10 even numbers")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}

	println()
	fmt.Println("for loop like a while loop. only with condition 1-10 odd numbers")
	i := 0

	for i < 10 {
		i++
		if i%2 == 0 {
			continue
		}
		print(i, " ")
	}
	println()

	i = 1
	println("unconditional loop with break. 1-10 multiples of 3")
	for {
		if i > 10 {
			break
		}

		if i%3 == 0 {
			print(i, " ")
		}
		i++
	}
	println()

	println("for loop not with initilizer but with conditon and post conditio")
	i = 1
	for ; i <= 10; i++ {
		if i%4 == 0 {
			print(i, " ")
		}
	}
	println()

	println("for loop not with initilizer  post condition but not conditon")
	for i := 1; ; i++ {
		if i > 10 {
			break
		}
		if i%4 == 0 {
			print(i, " ")
		}
	}
	println()
}
