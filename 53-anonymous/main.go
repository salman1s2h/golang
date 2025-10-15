package main

import "fmt"

func main() {

	var f1 Fn = func(i1, i2 int) int {
		return i1 + i2
	}
	r1 := f1(10, 20)
	fmt.Println(r1)

	s := f1.Sq(10)
	fmt.Println(s)

	map1 := make(map[string]any, 5)

	map1["greet"] = func() {
		fmt.Println("Hello World!")
	}
	map1["add"] = func(a, b int) int {
		return a + b
	}
	map1["mul"] = Mul
	map1["fn"] = f1
	map1["value"] = 100
	map1["done"] = true
	map1["something"] = struct{}{}

	println("\n--- iterating the map----")
	a, b := 10, 20
	for k, v := range map1 {

		switch v.(type) {
		case int:
			fmt.Println(k, "-->", v.(int))
		case bool:
			fmt.Println(k, "-->", v.(bool))
		case func():
			fmt.Println(k, "-->")
			v.(func())()
			fmt.Println("end of ", k)
			println()
		case func(a, b int) int:
			fmt.Println(k, "-->")
			c := v.(func(int, int) int)(a, b)
			fmt.Println("result-->", c)
			fmt.Println("end of ", k)
			println()
		case Fn:
			fmt.Println(k, "-->")
			c := v.(Fn)(a, b)
			fmt.Println("result-->", c)
			fmt.Println("end of ", k)
			println()
		default:
			fmt.Println(k, "--> type not found")
		}

	}

}

type Fn func(int, int) int

func (f Fn) Sq(a int) int {
	return a * a
}

func Mul(i, j int) int {
	return i * j
}
