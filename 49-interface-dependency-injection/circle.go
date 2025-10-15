package main

var Pi float32 = 3.14

type Circle float32

func (c Circle) Area() float32 {
	return Pi * float32(c*c)
}

func (c Circle) Perimeter() float32 {
	return 2 * Pi * float32(c)
}

func (c Circle) What() string {
	return "Circle"
}
