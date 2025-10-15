package main

import "fmt"

func main() {

	slice1 := make([]int, 3, 5)
	slice1[0], slice1[1], slice1[2] = 10, 20, 30
	//AddAndDouple(slice1, 40, 50, 60)
	slice1 = AddAndDoupleR(slice1, 40, 50, 60)
	fmt.Println(slice1)

	//slice2 := slice1 // header copy

	slice3 := make([]int, 2)
	//var slice3 []int
	copy(slice3, slice1) // desination slice must instantiated
	fmt.Println(slice3)

	slice4 := make([]int, 20)

	Copy(slice4, slice1)
	fmt.Println(slice4)
	clear(slice1) // make the slice as zero value slice
	fmt.Println(slice1)
}

func AddAndDouple(slice []int, nums ...int) {
	slice = append(slice, nums...)
	for i, v := range slice {
		slice[i] = v * 2
	}
}

func AddAndDoupleR(slice []int, nums ...int) []int {
	slice = append(slice, nums...)
	for i, v := range slice {
		slice[i] = v * 2
	}
	return slice
}
func Copy(s1, s2 []int) {
	if s1 == nil || s2 == nil {
		return
	}

	sl := min(len(s1), len(s2))

	for i := 0; i < sl; i++ {
		s1[i] = s2[i]
	}

}
