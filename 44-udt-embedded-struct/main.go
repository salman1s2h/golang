package main

import "fmt"

func main() {

	p1 := &Person{Name: "Jiten", Email: "Jiten@Outlook.Com", Contact: "9618558500", Status: "active", Address: &Address{City: "Bangalore", Pincode: "560086"}, SocialMedia: &SocialMedia{Social: map[string]string{"Linkedin": "linkedin.com/in/jpalaparthi"}}, Company: struct {
		Name    string
		WebSite string
	}{Name: "ICICI", WebSite: "icici.com"}, float32: 12312.123, E: struct{}{}}
	p1.Display()
}

type Person struct {
	Name        string
	Email       string
	Contact     string
	Status      string
	*Address                 // promoted field
	SocialMedia *SocialMedia // composition
	Company     struct {     // embedded struct
		Name    string
		WebSite string
	}
	float32          // anonymos field
	E       struct{} // empty struct field
}

func (p *Person) Display() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Contact:", p.Contact)
	p.Address.Display()
	p.SocialMedia.Display()
	fmt.Println("Company Name:", p.Company.Name)
	fmt.Println("Company Website:", p.Company.WebSite)
	fmt.Println("flaot32:", p.float32)
	fmt.Println("Empty struct E:", p.E)
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
