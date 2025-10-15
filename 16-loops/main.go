package main

import "math/rand"

func main() {

	for i, j := 1, 10; i <= j; i, j = i+1, j-1 {
		println("i:", i, "j:", j)
	}

	println("nested loop")

	done := false
	for i := 1; i <= 10; i++ {
		if done {
			break
		}
		for j := 2; j <= 10; j++ {
			if i == j {
				done = true
				break
			}
			println("i:", i, "j:", j)
		}
	}

out: // It is a label
	for i := 1; i <= 10; i++ {
		for j := 2; j <= 10; j++ {
			if i == j {
				break out
			}
			println("i:", i, "j:", j)
		}
	}

loop:
	for {
		num := rand.Intn(100)
		switch {
		case num%2 == 0:
			println(num, "is even")
		case num%2 == 1:
			println(num, "is odd")
			if num == 99 {
				break loop
			}
		}
	}

}
