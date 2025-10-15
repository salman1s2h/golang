package main

import "fmt"

func main() {
	var num1 int // on 64 bit machines it is 8 bytes on 32 bit 4 bytes
	// Zero value is zero for numbers
	var ok1 bool    // Zero value for bool --> false
	var str1 string // Zero value is ""

	println(num1, ok1, str1)
	num1, ok1, str1 = 100, true, "Hello World"
	println(num1, ok1, str1)

	var num2 int64 = 999999999999999999
	var num6 int64 = 999999999999999999
	var num7 int64 = num2 + num6 //this gives wrong result, have to use big int
	fmt.Println("Not a bignumber:", num7)
	var (
		num3 uint8  = 255
		num4 int32  = 12312312
		ok2  bool   = true
		str2 string = "Hello Universe"
	)

	fmt.Println(num2, num3, num4, ok2, str2)

	// Type inference --> The type is inferred by the compiler at compile time

	var num5 = 9999999 // int

	var float1 = 12312312.12312 // float64

	var ok3 = true

	var str3 = "Hello Customer, Good banking" // inferred as string

	fmt.Println("num5:", num5, "float1:", float1, "ok3:", ok3, "str3:", str3)
	fmt.Printf("num5: %d flaot1: %.3f ok3: %v str3: %s binary: %b hex: Ox%X\n", num5, float1, ok3, str3, num5, num5)
	// shorthand declaration

	a := 100
	b := 200
	t := a
	a = b
	b = t

	a, b = b, a // swapping in go
	fmt.Println("a,b", a, b)
	age := 39 // what is the problem with this approach
	var age1 uint8 = 39

	flaot2 := 123.34

	var float3 float32 = 123.34
	fmt.Println(age, age1, flaot2, float3)
	a1, b1, c1 := 10, 20, 30
	a1, b1, c1 = b1, c1, a1

	fmt.Println("a1,b1,c1", a1, b1, c1)

}

/* Numbers
int,int8,int16,int32,int64
uint, uint8,uint16,uint32,uint64
rune, byte (These are not special new datatypes but aliass)
float32,float64
uintptr
*/

// Boolean --> bool

// Strings --> string

// complex --> complex64,complex128

// interface --> any or interface{}
