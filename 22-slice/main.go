package main

import (
	"fmt"

	"math/rand"
)

func main() {
	//var slice2 []int
	slice1 := []int{10, 20, 30} // creating a slice with values
	if slice1 == nil {
		fmt.Println("slice is nil")
	} else {
		fmt.Println(slice1)
	}

	PrintHeader(slice1, "Slice1")

	fmt.Println("Sum of slcie1:", SumOf(slice1))

	slice2 := make([]int, 10) // short hand notation

	for i := range slice2 {
		slice2[i] = rand.Intn(500)
	}

	fmt.Println("Sum of slice2:", SumOf(slice2))

	arr1 := [5]int{10, 20, 30, 40, 50}

	SumOf(arr1[:]) // an array can be converted to slice

}

func PrintHeader(slice []int, name string) {
	fmt.Println(name, " header")
	fmt.Println("--------------")
	fmt.Printf("Slice header pointer:%p\n", &slice)
	if slice == nil {
		fmt.Println("Ptr:nil")
	} else {
		fmt.Printf("Ptr:%p\n", &slice[0])
	}
	fmt.Println("Len:", len(slice))
	fmt.Println("Cap:", cap(slice))
}

// make -> create a slice
// make --> create a map
// make --> channel

func SumOf(slice []int) int {
	sum := 0

	for _, v := range slice {
		sum += v
	}
	return sum
}
