package main

import "fmt"

func main() {
	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	slice1 = append(slice1, 0)
	slice2 := slice1 // Header {Ptr,Len,Cap}
	fmt.Printf("Address of slice1:%p\n", &slice1)
	fmt.Printf("Address of slice2:%p\n", &slice2)

	PrintHeader(slice1, "Slice1")
	println()
	PrintHeader(slice2, "Slice2")

	slice2[0] = 9999
	fmt.Println("slcie1", slice1)

	slice2 = append(slice2, 8888)
	slice2[1] = 1111
	fmt.Println("slcie1", slice1)
	fmt.Println("slcie2", slice2)
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
