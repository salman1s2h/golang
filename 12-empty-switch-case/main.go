package main

import "fmt"

func main() {
	num := -50

	switch { // empty switch
	case num >= 0 && num <= 50:
		fmt.Println(num, "is greater or 0 and less than or 50")
	case num > 50 && num <= 100:
		fmt.Println(num, "is greater 50 and less than or 100")
	case num > 100:
		fmt.Println(num, "is greater 100")
	default:
		fmt.Println(num, "is negative")
	}
}
