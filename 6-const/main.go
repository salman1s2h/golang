package main

import "fmt"

var G int = 100 // data segment

// const (
const PI float32 = 3.14 // constants are evaluated by the compiler and embedded in the code segment along with the code
// )
const (
	SECOND    = 1
	MINUTE    = 60 * SECOND
	HOUR      = 60 * MINUTE
	DAY       = 24 * HOUR
	SOMETHING = (100+200)/20 + (123-30)/MINUTE // not evaluated at runtime,evaluated at comp time
// can do expressions with constant, as along as all values or variables are constants
)

func main() {
	//PI = 3.14
	num1 := 30
	num2 := (100+200)/20 + (123-30)/MINUTE + num1 // This evaluated at runtime
	fmt.Println(num1, num2)

	const (
		MAX = 10
		MIN = 1
	)

	fmt.Println(num1, num2)
	fmt.Println(SECOND, MINUTE, HOUR, DAY, SOMETHING, MIN, MAX)
	// expression
	a, b, c, d := 10, 20, 30, 0
	d = a + b + c - (a+b)*c*10 + c/10 + (a * b * c)
	// c/10 is not c= c/10
	// d = a*b + a/b
	// d = a/b + a*b
	fmt.Println(d)
}

// Code/Text Segment
// Data Segment
// Stack Memory
// Heap memory
