package main

import (
	"fmt"

	"math/rand"
)

func main() {
	ishapes := make([]IShape, 0)
	ishapes = append(ishapes, Square(100.12), Circle(123.4), NewRect(12.4, 12.5), NewCuboid(13.23, 22.23, 23.2), NewRect(12.8, 16.7), NewCuboid(12.23, 23.23, 23.2), NewTriangle(12.3, 14.5, 13.5, 12))
	for _, v := range ishapes {
		Shape(v)
		println()
	}

	for i := 1; i <= 5; i++ {
		r := rand.Intn(7)
		Shape(ishapes[r])
		println()
	}

	r1 := NewRect(12.2, 13.4)
	a1 := r1.Area()
	p1 := r1.Perimeter()
	fmt.Println("Area of r1:", a1, "\nPerimeter of r1:", p1)

}

// itable --> interface definition table
//
