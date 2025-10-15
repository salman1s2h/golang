package main

import "fmt"

func main() {
	var slice1 []int // This is a declaration but not instantiation
	// The slice header has zero values and hence { Ptr: nil, Len: 0 and Cap:0}
	if slice1 == nil {
		println("yes slice1 is nil")
	}

	fmt.Printf("Actual Slice Header Pointer:%p\n", &slice1)

	PrintHeader(slice1, "slice1 declarearion time")
	slice1 = make([]int, 10) // instantiate
	// Slice1 {Ptr: 0x1123,Len:10, Cap:10}
	// What zero value of elements in the slice

	//fmt.Println(slice1)
	println()
	PrintHeader(slice1, "slice1")

	// var num int = 100
	// var ptr *int = &num
	// *ptr = 2000

	// var str string = "hello World"
	// var strPtr *string = &str
	// *strPtr = "hello universe"

	// var float1 float32 = 54.67
	// var floatPtr *float32 = &float1
	// *floatPtr = 123.234
}

// What is a slice --> Growable array
// Where does it allocate the memory --> Stack or Heap based on Go Compiler decision
// At runtime can I adde more elements to the slice--> Yes , can be appended
// When ever a slice is declared or created, a slice header is created
// Whad does the header contain?
// to instantiate a slice use a make built in function
/*
Slice Header
------------
Ptr
Len
Cap
*/

// str1:= "hello world" --> {Ptr: Len:}
// arr1 := [3]int{10,20,30} --> No header but the arr type itslef maintains the lenth .. it can be itereated based on the length and index
// slice1 :=[]int{10,20,30}--> {Ptr: Len: Cap:}

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
