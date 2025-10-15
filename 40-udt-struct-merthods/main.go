package main

import "fmt"

func main() {

	var r1 Rect = Rect{L: 123.3, B: 124.4}

	a1, p1 := Area(&r1), Perimeter(&r1)
	fmt.Println("Area of Rect r1:", a1)
	fmt.Println("perimeter of Rect r1:", p1)
	fmt.Println("Area A in Rect r1:", r1.A)
	fmt.Println("Perimeter P in Rect r1:", r1.P)

	r2 := Rect{L: 123.23, B: 234.34}
	a2 := r2.Area()
	p2 := r2.Perimeter()

	fmt.Println("Area of Rect r2:", a2)
	fmt.Println("perimeter of Rect r2:", p2)
	fmt.Println("Area A in Rect r2:", r2.A)
	fmt.Println("Perimeter P in Rect r2:", r2.P)

	r3 := NewRect(10.2, 12.4)
	a3 := r3.Area()
	p3 := r3.Perimeter()
	fmt.Println("Area of Rect r3:", a3)
	fmt.Println("perimeter of Rect r3:", p3)
	fmt.Println("Area A in Rect r3:", r3.A)
	fmt.Println("Perimeter P in Rect r2:", r3.P)

}

type Rect struct {
	L, B float32
	A    float32
	P    float32
}

// func Area(r Rect) float32 {
// r.A = r.L * r.B // no need to (*r).
// return r.A
// }
//
// func Perimeter(r Rect) float32 {
// r.P = 2 * (r.L + r.B)
// return r.P
// }

func Area(r *Rect) float32 {
	(*r).A = r.L * r.B // no need to (*r). rather use r. even it is a pointer
	return r.A
}

func Perimeter(r *Rect) float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func (r *Rect) Area() float32 {
	r.A = r.L * r.B // no need to (*r). rather use r. even it is a pointer
	return r.A
}

func (rect *Rect) Perimeter() float32 {
	rect.P = 2 * (rect.L + rect.B)
	return rect.P
}

// func Dummy1() (int, error) {
// return 0, nil
// }
//
// func Dummy2() (error, int) {
// return nil, 0
// }
//
