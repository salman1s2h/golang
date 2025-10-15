package main

import (
	"errors"
	"fmt"
)

func main() {
	slice1 := make([]int, 10, 15)
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4], slice1[5], slice1[6], slice1[7], slice1[8], slice1[9] = 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	//slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := slice1     // Header Copy
	slice3 := slice1[:5] // but not 5
	slice4 := slice1[2:8]
	slice5 := slice1[5:]
	PrintHeader(slice1, "slice1")
	PrintHeader(slice2, "slice2")
	PrintHeader(slice3, "slice3")
	PrintHeader(slice4, "slice4")
	PrintHeader(slice5, "slice5")
	slice6 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if slice6, err := DeleteElem(slice6, 9); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(slice6)
	}

	//slice6 = append(slice6[:2], slice6[3:]...)
	if slice6, err := DeleteElem(slice6, 2); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(slice6)
	}

}

func DeleteElem(slice []int, i uint) ([]int, error) {
	if slice == nil {
		return nil, errors.New("nil slice")
	}

	if i >= uint(len(slice)) {
		return nil, errors.New("invalid index")
	}
	if int(i) == len(slice)-1 {
		return slice[:len(slice)-1], nil
	}
	slice = append(slice[:i], slice[i+1:]...)

	return slice, nil
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
