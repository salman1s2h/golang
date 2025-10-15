package main

import "fmt"

func main() {

	c1 := ColourCode{int: 100, string: "Red", mystring: "Dark Red"}
	fmt.Println(c1)

}

type mystring = string

// struct with anonymous fields
type ColourCode struct {
	int      // anonymous fields
	string   // anonymous fields
	mystring // anonymous fields
}
