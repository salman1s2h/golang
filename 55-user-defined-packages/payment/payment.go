package payment

import (
	"errors"
	"fmt"
)

// UpiPayment to use camel case or pascal case depends on accessibility
// upiPayment

// upi_payment // not a preffered to use snake case

var g int // unexported | restricted

var G int // exported | unrestricted

func greet() {
	fmt.Println("hello payment package")
}

func Greet() {
	greet()
}

type upiPayment struct {
	AccountNo string
	Amount    float32
	Receiver  string
	status    string // unexported or restricted
}

type Payment struct {
	AccountNo string
	Amount    float32
	Receiver  string
	status    string // unexported or restricted
}

func (p *Payment) Transfer() error {
	if p == nil {
		return errors.New("nil Payment object")
	}
	fmt.Println(p, p.Amount, "Transffered to :", p.Receiver)
	return nil
}

func (p *Payment) display() {
	fmt.Println(p)
}

type IPayment interface {
	Transfer() error
}
