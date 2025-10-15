package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var e1 E

	fmt.Println("Size of e1:", unsafe.Sizeof(e1))

	e1.Greet()

	var e2, e3 E
	var ptrE1 = &e1
	var ptrE2 = &e2

	fmt.Printf("%p\n", ptrE1)
	fmt.Printf("%p\n", ptrE2)
	fmt.Printf("%p\n", &e3)
	num := new(int)
	fmt.Printf("%p\n", num)

}

type E struct{}

func (e *E) Greet() {
	fmt.Println("Hello World")
}

type T struct{}

func (e *T) Greet() {
	fmt.Println("Hello Universe")
}
