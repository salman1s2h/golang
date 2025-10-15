package main

func main() {
	var ptr *int = new(int) // 8 bytes is allpocated and assign a zero value
	println(*ptr)
	println(ptr)
	*ptr = 100
	println(*ptr)

	str1 := new(string) // How much memory is allocated here?
	if str1 == nil {
		println("str1 yes nil")
	} else {
		println("str1 not nil")
	}
	var str2 *string
	if str2 == nil {
		println("str2 yes nil")
	} else {
		println("str2 not nil")
	}

}
