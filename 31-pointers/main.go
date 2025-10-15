package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr1 := [5]int{10, 11, 12, 13, 14}

	// var ptr uintptr = 1374390755328+8
	// ptr += 8

	// fmt.Printf("address of 0th %p %d", &arr1[0], &arr1[0])
	// &arr1[0] += 8
	// unsafe.Pointer
	ptr1 := uintptr(unsafe.Pointer(&arr1[0])) // 1 and 4
	ptr1 += unsafe.Sizeof(arr1[0])            // arithmetic operation on uintptr variable
	val := *(*int)(unsafe.Pointer(ptr1))      // 3 ane 2
	fmt.Println(val)

	ptr2 := uintptr(unsafe.Pointer(&arr1[0]))
	for i := 0; i < len(arr1); i++ {
		val := *(*int)(unsafe.Pointer(ptr2))
		fmt.Println(val)
		ptr2 += unsafe.Sizeof(arr1[0])
	}

}

// 1. A pointer value of any type can be converted to a Pointer.
// 2. A Pointer can be converted to a pointer value of any type.
// 3. A uintptr can be converted to a Pointer.
// 4. A Pointer can be converted to a uintptr.
