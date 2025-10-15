package main

import (
	//"fmt"
	"fmt"
	"math/rand"

	"github.com/JitenPalaparthi/icici-batch4-shapes/circle"
	"github.com/JitenPalaparthi/icici-batch4-shapes/cuboid"
	"github.com/JitenPalaparthi/icici-batch4-shapes/demo"

	. "time"

	rectangle "github.com/JitenPalaparthi/icici-batch4-shapes/rect"
	. "github.com/JitenPalaparthi/icici-batch4-shapes/shape"
	"github.com/JitenPalaparthi/icici-batch4-shapes/square"
	"github.com/JitenPalaparthi/icici-batch4-shapes/triangle"
	_ "github.com/gin-gonic/gin"
)

func main() {
	Greet()
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
}

// itable --> interface definition table
//
