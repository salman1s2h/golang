package main

import (
	"fmt"
	"unsafe"
)

func main() {

	str1 := "Hello World!"

	strPtr := (*[2]int)((unsafe.Pointer(&str1)))
	fmt.Println(strPtr[0])
	strPtr[1] = 100
	fmt.Println(strPtr[1])
	fmt.Println(len(str1))

	slice1 := make([]int, 3, 10)
	slice1[0] = 10
	slice1[1] = 11
	slice1[2] = 12
	sliceheader := (*[3]int)((unsafe.Pointer(&slice1)))
	fmt.Println(sliceheader[0])
	fmt.Println(sliceheader[1])
	fmt.Println(sliceheader[2])
	sliceheader[1] = 20
	sliceheader[2] = 5000

	fmt.Println("len:", len(slice1), "cap:", cap(slice1))

	// for i := 1; i <= 4995; i++ {
	// slice1 = append(slice1, i)
	// }

	for i, v := range slice1 {
		fmt.Println(i, "-->", v)
	}

}

// String header
// Ptr
// Len
