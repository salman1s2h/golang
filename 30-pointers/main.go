package main

import "fmt"

func main() {
	var num1 int = 100
	var ptr1 *int

	if ptr1 == nil {
		println("yes ptr is nil", ptr1)
	}
	ptr1 = &num1

	*ptr1 = 200 // dereferencing
	//
	// var ptr2 *int
	//
	// *ptr2 = 100

	Incr(&num1)

	fmt.Println(num1)

	arr1 := [3]int{10, 20, 30}

	ptr2 := &arr1[0]
	//ptr1 += 8 // pointer arithmetic
	*ptr2 = 100

	fmt.Println(arr1)

	a, b, c := 10, 20, 30

	arrPtr := [3]*int{&a, &b, &c}
	*arrPtr[0] = 100
	*arrPtr[1] = 200
	*arrPtr[2] = 300

	println(a, b, c)

	var ptrArr *[3]int = &arr1
	//var ptrArr1 *[3]*int = &arrPtr
	fmt.Println((*ptrArr)[0])
	fmt.Println(ptrArr[1])

	for i, v := range ptrArr {
		fmt.Println(i, "---", v)
	}

	var ptr3 **int = &ptr1
	var ptr4 ***int = &ptr3

	***ptr4 = 99999

	println(num1)

}

// What does pointer hold
// What is the zero value of a pointer
// To derefernce pointer should not nil

// nil pointer dereferencing
// double free (not in Go)
// dangling pointer

// func Sq(i int) *int {
// ptr := &i
// *ptr = i * i
// return ptr
// }

func Incr(ptr *int) { // call/pass by reference
	if ptr != nil {
		*ptr++
	}
}
