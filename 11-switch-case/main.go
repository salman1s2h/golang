package main

import "fmt"

func main() {

	day := uint8(120)

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("Noday")
	}

	switch day {
	case 2, 3, 4, 5, 6:
		fmt.Println("Working day")
	case 1, 7:
		fmt.Println("Weekend Holiday")
	default:
		fmt.Println("invalid day")
	}
}
