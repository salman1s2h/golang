package main

import "fmt"

func main() {
	func( /*inpu params*/ ) /*return types*/ {
		//body
	}( /*executor*/ )

	func() {
		println("Hello World")
	}()

	c1 := func(a, b int) int {
		return a + b
	}(10, 20)

	fmt.Println(c1)
	num := 100
	f1 := func(a, b int) int {
		return (a - b) / num
	}
	if f1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("not nil")
	}

	c2 := f1(10, 20)
	println(c2)

	c3 := Calc(100, 20, func(i1, i2 int) int {
		return i1 / i2
	})
	println(c3)

	c4 := Calc(10, 3, Mul)
	println(c4)

	c5 := Calc(10, 20, f1)
	println(c5)

}

func Calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func Mul(i, j int) int {
	return i * j
}
