package main

import "fmt"

func main() {
	slice1 := make([]int, 10)
	//[0 0 0 0 0 0 0 0 0 0]
	slice1 = append(slice1, 9999)
	//[0 0 0 0 0 0 0 0 0 0 9999]
	fmt.Println(slice1)
	//slcie1[0] = 9999

	var slice2 []int

	slice2 = append(slice2, 999) // When slice is nill append instantiate the slice and appends the value

	var slice3 []int

	for _, v := range slice3 { // even though the slice is nil , the range funcion works as the len is 0
		fmt.Println(v)
	}
}

// append can be used on nil slice as well
// range loop can be used even the slice is nil but does not iterate since the len in the header is 0
