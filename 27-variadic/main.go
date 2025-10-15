package main

import (
	"fmt"

	"math/rand"
)

func main() {

	//slice1 := append([]int{}, 10, 20, 30, 40, 50, 60)
	//fmt.Println("Hello", "World", true, (10 + 10), (10 / 10), (true && true || false))
	fmt.Println(SumOf())
	fmt.Println(SumOf(10, 20))
	fmt.Println(SumOf(10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 140, 150))
	slice1 := make([]int, 100000)
	for i := range slice1 {
		slice1[i] = rand.Intn(10000)
	}
	// a slice can be passed as ar argument to a variadic parameters

	fmt.Println("Slice1", SumOf(slice1...))

	slice2 := MulOf(2, 1, 2, 3, 4, 5)
	fmt.Println(slice2)

	arr1 := [...]int{10, 20, 30, 40, 50, 60}
	fmt.Println("SumOf arr1:", SumOf(arr1[:]...))

}

// Variadic parameter should be the last parameter..
// Variadic parameter cannot be used like a data type but can only be used in functions and methods

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func MulOf(nums ...int, num int) []int { //variadic must be the last param
func MulOf(num int, nums ...int) []int {
	slice := make([]int, len(nums))
	for i, v := range nums {
		slice[i] = v * num
	}
	return slice
}
