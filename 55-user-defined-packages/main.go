package main

import (
	//"fmt"
	"fmt"
	"math/rand"
	"shapes/circle"
	"shapes/cuboid"
	"shapes/demo"
	pay "shapes/payment"

	rectangle "shapes/rect"
	. "shapes/shape"
	"shapes/square"
	"shapes/triangle"
	. "time"
)

func main() {
	ishapes := make([]IShape, 0)
	ishapes = append(ishapes, square.Square(100.12), circle.Circle(123.4), rectangle.NewRect(12.4, 12.5), cuboid.NewCuboid(13.23, 22.23, 23.2), rectangle.NewRect(12.8, 16.7), cuboid.NewCuboid(12.23, 23.23, 23.2), triangle.NewTriangle(12.3, 14.5, 13.5, 12))
	for _, v := range ishapes {
		Shape(v)
		println()
	}

	for i := 1; i <= 5; i++ {
		r := rand.Intn(7)
		Shape(ishapes[r])
		println()
	}

	r1 := rectangle.NewRect(12.2, 13.4)
	a1 := r1.Area()
	p1 := r1.Perimeter()
	fmt.Println("Area of r1:", a1, "\nPerimeter of r1:", p1)
	fmt.Println(Now())
	demo.Now()

	// export or not

	p := pay.Payment{AccountNo: "101", Amount: 12312.23, Receiver: "909", status: "active"}
	//payment.greet() // cant call this since it is not exported
	pay.Greet() // can call this since it is exported
	p.Transfer()

	//u := pay.upiPayment{}
}

// itable --> interface definition table
//
