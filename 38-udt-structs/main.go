package main

import "fmt"

func main() {
	var p1 Person     // created it
	p1.Name = "Jiten" // accessing fields
	p1.Email = "Jitenp@Outlook.Com"
	p1.Contact = "9618558500"

	p2 := Person{Name: "John", Email: "John@Gmail.Com", Contact: "91231223"}

	var p3 Person = Person{Name: "John", Email: "John@Gmail.Com", Contact: "91231223"}

	fmt.Println(p1)
	fmt.Println(p2)
	println()

	Display(p1)
	println()
	Display(p2)
	println()

	Display(p3)
}

// type Student = Person
// type rune = int32
// type any = interface{}

type Person struct {
	Name    string
	Email   string
	Contact string
}

func Display(p Person) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Contact:", p.Contact)
}
