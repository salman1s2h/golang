package main

import "fmt"

func main() {
	defer func() {
		defer fmt.Println("end of  func-1")
		defer println("End of main")
		println()
		fmt.Println("calling func-1")
	}()
	println("Hello World")
	println("Start of main")

	str := "Hello World"

	for _, v := range str {
		defer print(string(v), " ")
	}

	// <= 1.20 problem with loop variable
	for i := 0; i < len(str); i++ {
		// defer func(j int) {
		// println(string(str[j]))
		// }(i)
		defer println(string(str[i]))
	}

	i := 0
loop:
	if i < len(str) {
		defer print(" ", string(str[i]))
		i++
		goto loop
	}

	funcslice := make([]func(), 0)

	j := 0
loop1:
	if j < 10 {
		funcslice = append(funcslice, func() {
			println("calling a funcSlice func", j)
		})
		j++
		goto loop1
	}

	for _, v := range funcslice {
		v()
	}

}
