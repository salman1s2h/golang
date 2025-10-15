package main

import (
	"fmt"

	"math/rand"
)

func main() {
	slice1 := make([]int, 10) // short hand notation

	for i := range slice1 {
		slice1[i] = rand.Intn(500)
	}

	PrintHeader(slice1, "Slice1 before append")
	slice1 = append(slice1, 9999)
	println()
	PrintHeader(slice1, "Slice1 after append")

	println()
	slice2 := make([]int, 10, 20) // short hand notation
	for i := range slice2 {
		slice2[i] = rand.Intn(500)
	}
	PrintHeader(slice2, "slice2 before append")
	slice2 = append(slice2, 9999)
	println()
	PrintHeader(slice2, "slice2 after append")

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
	fmt.Println(name, ":", slice)
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
