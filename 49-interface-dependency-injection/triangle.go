package main

func NewTriangle(a, b, c, h float32) *Triangle {
	return &Triangle{A: a, B: b, C: c, H: h}
}

type Triangle struct {
	A, B, C, H float32
}

func (t *Triangle) Area() float32 {
	return .5 * t.B * t.H
}

func (t *Triangle) Perimeter() float32 {
	return t.A + t.B + t.C
}

func (t *Triangle) What() string {
	return "Triangle"
}
