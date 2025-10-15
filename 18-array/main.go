package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//var num int = 100

	//const pi float32 = 3.14

	var arr1 [50]int
	var arr2 [4]int
	arr2[0], arr2[1], arr2[2], arr2[3] = 10, 20, 30, 40

	// if arr2 == arr1 { // cannpot perform == operation bcz both of them are differnet

	// }
	// zero value of an array [0 0 0 0 0]

	arr1[0] = 100
	arr1[1] = 1234
	arr1[2] = 123
	for i, v := range arr1 {
		if v == 0 {
			arr1[i] = rand.Intn(100)
		}
	}

	fmt.Println(arr1)

	fmt.Println("Sumof", SumOf(arr1))

	//SumOf(arr2) // this does not work due to input type is differnet than the argument
	SumOf4(arr2)
}

// Arrays are fixed size
// The length to be known to the compiler
// What is the type of arr1-> the type of the array includes its length
// cannot type cast from one type of the array to another type
// What is zero value of arr1

func SumOf(arr [50]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumOf4(arr [4]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
