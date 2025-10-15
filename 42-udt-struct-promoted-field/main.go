package main

import "fmt"

func main() {

	p1 := Person{Name: "Jiten", Email: "Jiten@Outlook.Com", Contact: "9618558500", Status: "active", Address: &Address{City: "Bangalore", Pincode: "560086", Status: "available"}, SocialMedia: SocialMedia{Social: map[string]string{"linkedin": "linkedin.com/in/jpalaparthi", "x": "x.com/jitenp"}}}
	p1.Display()
	fmt.Println(p1.Address.City)
	fmt.Println(p1.City)
	fmt.Println(p1.Social)
	p1.Display()
	p1.Address.Display()
	p1.SocialMedia.Display()
	p1.DisplayA()

}

type Person struct {
	Name        string
	Email       string
	Contact     string
	Status      string
	*Address    // promoted field
	SocialMedia // promoted field
}

func (p *Person) Display() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Contact:", p.Contact)
	p.Address.Display()
	p.SocialMedia.Display()
}

type Address struct {
	City, Pincode, Status string
}

func (a *Address) Display() {
	fmt.Println("City:", a.City)
	fmt.Println("Pincode:", a.Pincode)
}

func (a *Address) DisplayA() {
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
