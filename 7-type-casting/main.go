package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num1 uint8 = 100
	var num2 int32 = int32(num1)
	println(num1, num2)

	var float1 float32 = 123.23
	var num3 int64 = int64(float1)

	fmt.Println(num3)

	//var ok1 bool = true

	//var num4 uint8 = uint8(ok1) // not possible

	var num4 uint64 = 123123123
	var num5 uint8 = uint8(num4)

	fmt.Printf("binary of num4: %b decimal %d\n", num4, num4)
	fmt.Printf("binary of num5: %b decimal %d\n", num5, num5)

	var num6 uint16 = uint16(num4) // 1011010110110011 //46515
	fmt.Println(num6)

	var num7 = 19000

	var str1 = string(num7) // byte to char casting..

	fmt.Println(str1)

	// as it is
	str2 := fmt.Sprint(num7)
	fmt.Println(str2)

	str3 := strconv.Itoa(num7)
	fmt.Println(str3)

	// v := Calc(10, 30)
	// fmt.Println(v)

	a1, s1, m1, d1 := Calc(30, 40)
	fmt.Println(a1, s1, m1, d1)

	a2, s2, m2, _ := Calc(30, 40)
	fmt.Println(a2, s2, m2)

	_, _, _, d2 := Calc(30, 40)
	fmt.Println(d2)

	_, s3, _, d3 := Calc(10, 40)
	fmt.Println(s3, d3)

	var str4 string = "112323123"
	v1, err := strconv.Atoi(str4) // may or mannot
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Str to int:", v1)
	}

	var str5 string = "112a23123"
	v1, err = strconv.Atoi(str5) // may
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Str to int:", v1)
	}

	str5 = "1231123.1231"

	f1, err := strconv.ParseFloat(str5, 32)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Str to float:", f1)
	}

	str4 = "0"

	b1, err := strconv.ParseBool(str4)
	fmt.Println(b1, err)

	var ok1 bool = false
	num8 := uint8(BoolToInt(ok1))
	fmt.Println(num8)

}

// type casting --> no implicit type casting in Go
// type conversion
// type assertion

// func Calc(a, b int) int {
// return a + b
// }
//

func Calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
