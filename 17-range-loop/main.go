package main

func main() {

	str1 := "Hello World"
	str2 := "Hello, 世界!"

	for i := 0; i < len(str1); i++ {
		print(str1[i], "---->", string(str1[i]), " ")
	}
	println()

	for i := 0; i < len(str2); i++ {
		print(str2[i], "---->", string(str2[i]), " ")
	}
	println()
	println("Using range loop")

	for _, v := range str1 {
		println(string(v))
	}

	for i := range str1 {
		println("index:", i)
	}

	for i, v := range str2 {
		println(i, string(v))
	}
	println()

	// 1.22 onwards
	for i := range 10 {
		if i%2 == 0 {
			print(i, " ")
		}
	}

	// a, _ := calc(10, 20)
	// println(a)
	// _, b := calc(10, 20)
	// println(b)

}

func calc(i, j int) (int, int) {
	return i + j, i - j
}

// range loop on strings, array ,slice --> it returns index and value
// range loopm on map , it returns key and value
// range loop on a channel returns a value from the channel
// range of values
