package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		m    int     = 100
		m1   Myint1  = 200
		m2   Myint2  = 120
		m3   Myint3  = 150
		f1   float64 = 123.234
		any1 any     = 100
	)

	fmt.Println("m1 ToString:", Myint1(m).ToString())
	fmt.Println("Sq on m", Myint2(m).Sq())
	fmt.Println("Cube on m", Myint3(m).Cube())
	fmt.Println("Sqrt on m", Myint4(m).Sqrt())

	println()
	fmt.Println("m1 ToString:", m1.ToString())
	fmt.Println("Sq on m1", Myint2(m1).Sq())
	fmt.Println("Cube on m1", Myint3(m1).Cube())
	fmt.Println("Sqrt on m1", Myint4(m1).Sqrt())

	println()
	fmt.Println("m2 ToString:", Myint1(m2).ToString())
	fmt.Println("Sq on m2", m2.Sq())
	fmt.Println("Cube on m2", Myint3(m2).Cube())
	fmt.Println("Sqrt on m2", Myint4(m2).Sqrt())

	println()
	fmt.Println("m3 ToString:", Myint1(m3).ToString())
	fmt.Println("Sq on m3", Myint2(m3).Sq())
	fmt.Println("Cube on m3", m3.Cube())
	fmt.Println("Sqrt on m3", Myint4(m3).Sqrt())

	println()
	fmt.Println("f1 ToString:", Myint1(f1).ToString())
	fmt.Println("Sq on f1", Myint2(f1).Sq())
	fmt.Println("Cube on f1", Myint3(f1).Cube())
	fmt.Println("Sqrt on f1", Myint4(f1).Sqrt())

	println()
	fmt.Println("any1 ToString:", Myint1(any1.(int)).ToString())
	fmt.Println("Sq on any1", Myint2(any1.(int)).Sq())
	fmt.Println("Cube on any1", Myint3(any1.(int)).Cube())
	fmt.Println("Sqrt on any1", Myint4(any1.(int)).Sqrt())

}

// int
type Myint1 int
type Myint2 int
type Myint3 int

type Myint4 Myint3

func (m Myint1) ToString() string {
	return fmt.Sprint(m)
}

func (m Myint2) Sq() int {
	return int(m * m)
}

func (m Myint3) Cube() int {
	return int(m * m * m)
}

func (m Myint4) Sqrt() float64 {
	return math.Sqrt(float64(m))
}
