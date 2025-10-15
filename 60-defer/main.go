package main

import "fmt"

func main() {
	v := 100
	defer func(v int) {
		fmt.Println("deffered", v)
	}(v)

	v += 1
	fmt.Println("sequencial", v)

	// a := 1
	// {
	// a := 1
	// println(a)
	// }
	// println(a)

}
