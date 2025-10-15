package main

import "fmt"

func main() {
	var any1 interface{}
	any1 = 100.5
	//var num1 int = any1.(int) // any.(type)
	//fmt.Println(num1)

	var num1 int
	num1, done := any1.(int)
	if done {
		fmt.Println(num1)
	} else {
		fmt.Println("could not be asserted to int so trying to with float64")
		float1, done := any1.(float64)
		if done {
			fmt.Println(float1)
		} else {
			fmt.Println("Could not assert to int")
		}
	}
}
