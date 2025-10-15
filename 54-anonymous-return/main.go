package main

import "fmt"

func main() {
	f1 := Add(10, 20)
	r1 := f1()
	fmt.Println(r1)
}

// func return another function
func Add(a, b int) func() int {
	return func() int {
		return a + b
	}
}
