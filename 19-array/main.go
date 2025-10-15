package main

import "fmt"

func main() {

	arr1 := [4]int{10, 20, 30, 40} // short hand declaration

	arr2 := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} // the length is evaluated by the compiler
	fmt.Println(arr1, arr2)
	arr2d := [2][2]int{{1, 2}, {3, 4}}
	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	//arr4d := [1][2][2][2]int{{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}}

	for _, v1 := range arr2d {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		println()
	}
	println()
	for _, v1 := range arr3d {
		for _, v2 := range v1 {
			for _, v3 := range v2 {
				print(v3, " ")
			}
		}
		println()
	}

	println()

	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			fmt.Print(arr2d[i][j], " ")
		}
		println()
	}

}
