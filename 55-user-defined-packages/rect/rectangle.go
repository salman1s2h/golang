package rectangle

import "fmt"

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

type Rect struct {
	L, B float32
}

func (r *Rect) Display() {
	fmt.Println("L:", r.L)
	fmt.Println("B:", r.B)
}

func (s *Rect) What() string {
	return "Rect"
}
