package main

func main() {
	var num1, num2 int = 100, 200
	println(num1)
	//fmt.Println(num1)
	//fmt.Printf("address of num1 %p\n", &num1)
	println(num2)
	ptr := &num2
	println(ptr)

	sq1 := Sq(num2)
	println(sq1)

	sq2 := SqP(num2)
	println(*sq2)
	ptr1 := new(int)
	var num3 = 10
	SqPI(num3, ptr1)
	println(*ptr1)

	slice1 := make([]int, 1000000)
	println(len(slice1))

	slice2 := make([]int, 1000)
	println(len(slice2))
}

func Sq(i int) int {
	return i * i
}

func SqP(i int) *int {
	p := new(int) //
	*p = i * i
	return p // Dangling pointer
}

func SqPI(i int, p *int) {
	if p != nil {
		*p = i * i
	}
}
