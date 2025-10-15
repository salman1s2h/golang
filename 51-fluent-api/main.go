package main

import "fmt"

func main() {

	val := NewCalc(10).Add(2).Add(-2).Sub(1).Mul(4).Add(10).Sub(10).Div(2).Get()
	fmt.Println(val)

	// builder pattern

	r1 := &Rect{}
	r1.SetRect(10.2, 13.4)
	r1.Display()
	r1.SetColour("red")
	r1.Display()

	println()
	r2 := &Rect{}
	//(&Rect{}).SetRect(12.12, 12.5).SetColour("Blue").Display()
	r2.SetRect(12.12, 12.5).SetColour("Blue").Display()
	println()
	r2.SetBgColour("Grey").Display()

}

type ICalc interface {
	Add(n int) ICalc
	Sub(n int) ICalc
	Mul(n int) ICalc
	Div(n int) ICalc
	Get() int
}

type Calc struct {
	data int
}

func NewCalc(d int) *Calc {
	return &Calc{data: d}
}

func (c *Calc) Add(n int) ICalc {
	c.data += n
	return c
}

func (c *Calc) Sub(n int) ICalc {
	c.data -= n
	return c
}

func (c *Calc) Mul(n int) ICalc {
	c.data *= n
	return c
}

func (c *Calc) Div(n int) ICalc {
	c.data /= n
	return c
}

func (c *Calc) Get() int {
	return c.data
}

type IShape interface {
	SetRect(l, b float32) IShape
	SetColour(c string) IShape
	SetBgColour(c string) IShape
	Display()
}

type Rect struct {
	L, B     float32
	Colour   string
	BgColour string
}

func (r *Rect) SetRect(l, b float32) IShape {
	r.L = l
	r.B = b
	return r
}

func (r *Rect) SetColour(c string) IShape {
	r.Colour = c
	return r
}

func (r *Rect) SetBgColour(c string) IShape {
	r.BgColour = c
	return r
}

func (r *Rect) Display() {
	fmt.Println("L:", r.L)
	fmt.Println("B:", r.L)
	fmt.Println("Colour:", r.Colour)
	fmt.Println("Bg Colour:", r.BgColour)
}
