package main

import "fmt"

func main() {

	var p1 struct {
		Name    string
		Email   string
		Contact string
	}

	p1.Contact = "9618558500"
	p1.Email = "JitenP@Outlook.Com"
	p1.Contact = "9618558500"

	var p2 struct {
		Name    string
		Email   string
		Contact string
	} = struct {
		Name    string
		Email   string
		Contact string
	}{Name: "Jiten", Email: "Jitenp@outlook.com", Contact: "9618558500"}

	fmt.Println(p1)
	println()
	fmt.Println(p2)

	Display(p1)
	Display(p2)

	var p3 struct {
		Name    string
		Email   string
		Contact string
		Status  string
	} = struct {
		Name    string
		Email   string
		Contact string
		Status  string
	}{Name: "Jiten", Email: "Jitenp@outlook.com", Contact: "9618558500", Status: "active"}
	//Display(p3) cannot use Display since the input prameter does not match with p3 variable'
	println()
	fmt.Println(p3)
}

func Display(p struct{ Name, Email, Contact string }) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Contact:", p.Contact)

}
