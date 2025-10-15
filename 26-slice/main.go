package main

import "fmt"

func main() {

	slice1 := []int{}        // Ptr:Some Dummy/PlaceHolder Pointer  Len:0 Cap:0
	slice2 := make([]int, 0) // Ptr:Some Dummy/PlaceHolder Pointer  Len:0 Cap:0
	var slice3 []int

	if slice1 == nil {
		fmt.Println("slice1 is nil")
	}

	if slice2 == nil {
		fmt.Println("slice2 is nil")
	}

	if slice3 == nil {
		fmt.Println("slice3 is nil")
	}

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
