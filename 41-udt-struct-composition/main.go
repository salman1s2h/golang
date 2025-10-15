package main

import "fmt"

func main() {

	p1 := Person{Name: "Jiten", Email: "Jiten@Outlook.Com", Contact: "9618558500", Address: &Address{City: "Bangalore", Pincode: "560086"}, SocialMedia: SocialMedia{Social: map[string]string{"linkedin": "linkedin.com/in/jpalaparthi", "x": "x.com/jitenp"}}}
	p1.Display()
}

type Person struct {
	Name        string
	Email       string
	Contact     string
	Address     *Address    // composition
	SocialMedia SocialMedia // composition
}

func (p *Person) Display() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Contact:", p.Contact)
	p.Address.Display()
	p.SocialMedia.Display()
}

type Address struct {
	City, Pincode string
}

func (a *Address) Display() {
	fmt.Println("City:", a.City)
	fmt.Println("Pincode:", a.Pincode)
}

type SocialMedia struct {
	Social map[string]string
}

func (s *SocialMedia) Display() {
	for k, v := range s.Social {
		fmt.Println(k, "->", v)
	}
}
