package main

import "fmt"

type IShape interface {
	Area() float32
	Perimeter() float32
	IWhat
}

type IWhat interface {
	What() string
}

func Shape(ishape IShape) {
	a1 := ishape.Area()
	p1 := ishape.Perimeter()
	fmt.Println("Area of ", ishape.What(), "a1:", a1)
	fmt.Println("Perimeter of ", ishape.What(), "p1:", p1)
}
