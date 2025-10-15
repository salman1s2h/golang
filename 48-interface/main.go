package main

import (
	"fmt"
)

func main() {
	var ishape IShape
	//var any1 interface{}
	//fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))
	ishape = Square(12.4)
	//any1 = 100
	//var iempty IEmpty = 100
	a1 := ishape.Area()
	p1 := ishape.Perimeter()

	fmt.Println("Area of ", ishape.What(), "a1:", a1)
	fmt.Println("Perimeter of ", ishape.What(), "p1:", p1)

	r1 := NewRect(10.4, 15.6)

	ishape = r1 // what is DataPtr and what is TypePtr
	a1 = ishape.Area()
	p1 = ishape.Perimeter()
	fmt.Println("Area of ", ishape.What(), "a1:", a1)
	fmt.Println("Perimeter of ", ishape.What(), "p1:", p1)

}

type IEmpty interface {
}

type IShape interface {
	Area() float32
	Perimeter() float32
	IWhat
}

type IWhat interface {
	What() string
}

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return float32(4 * s)
}
func (s Square) What() string {
	return "Square"
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

type Rect struct {
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

func (r *Rect) Display() {
	fmt.Println("L:", r.L)
	fmt.Println("B:", r.B)
}

func (s *Rect) What() string {
	return "Rect"
}
