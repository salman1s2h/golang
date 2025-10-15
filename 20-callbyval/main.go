package main

func main() {
	num := 100
	Incr(num) // call by or pass by value
	println(num)
}

func Incr(n int) { // call by value and pass by reference
	n++
}
