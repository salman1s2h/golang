package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age uint8 = 16
	if age >= 18 {
		println("eligible to vote", age)
	} else {
		println("not eligible to vote ", age)
	}

	if num, err := strconv.Atoi("213123123"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num)
	}

	fmt.Println('f', 'M')
	age1, gender1 := uint8(18), 'M'
	//fmt.Println(age1 >= 18 && gender1 == 'F' || gender1 == 102 && true || (true || false))
	if age1 >= 18 && (gender1 == 'F' || gender1 == 102) {
		fmt.Println("She is eligible for marraiage becase of age is", age1)
	} else if age1 >= 21 && (gender1 == 77 || gender1 == 'm') {
		fmt.Println("He is eligible for marraiage becase of age is", age1)
	} else {
		fmt.Println("not eligible for marriage")
	}

}
