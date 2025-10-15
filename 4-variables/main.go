package main

import "fmt"

func main() {
	var char1 int32 = 'A'
	var char2 uint8 = 'B'
	var char3 uint16 = 'C'
	var char4 int64 = 'D'
	var char5 rune = 'E'
	var char6 int32 = '世'
	var char7 rune = '界'
	var char9 float32 = 'A'

	var char8 rune = char6 + 5

	fmt.Println(char1, char2, char3, char4, char5, char6, char7, char8, char9)
	//fmt.Println(string(char6), string(char8))

	// complex

	c1 := 123.23 + 1.2i        // shorthand declaration
	c2 := complex(123.23, 1.2) // creating a complex number using builtin function

	var r1 float32 = 123.23
	var im1 float32 = 1.2

	c3 := complex(r1, im1)

	fmt.Println(c1, c2, c3)

	r2, im2 := real(c1), imag(c1)
	fmt.Println("Real part:", r2, "Imag part:", im2)

}
