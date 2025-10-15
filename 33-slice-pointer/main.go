package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 10) // cap=:10 len:10

	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}
	fmt.Println(slice1)
	AddAndDouple(&slice1, 99, 88, 77, 66, 55, 44, 33, 22, 11)
	fmt.Println(slice1)
}

func AddAndDouple(slice *[]int, nums ...int) {
	if slice == nil {
		return
	}
	*slice = append(*slice, nums...)
	for i, v := range *slice {
		(*slice)[i] = v * 2
	}
}
