package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//type iface = interface{}

func main() {

	var num1 int

	fmt.Println("Val of num1:", num1)
	fmt.Println("Type of num1:", reflect.TypeOf(num1), "Size of:", unsafe.Sizeof(num1))

	var any1 any // what is the zero value of any
	//var any2 iface
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	if any1 == nil {
		println("yes any1 is nil")
	}

	any1 = 100
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	any1 = true
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	any1 = 'A'
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	any1 = "Hello World!"
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

	any1 = 123121.12312
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1), "Size of any1:", unsafe.Sizeof(any1))

}

/*
any Header
----------
DataPtr * Zero Value = nil
TypePtr * Zero Value = nil
*/
