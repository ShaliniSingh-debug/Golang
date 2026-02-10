package main

import (
	"fmt"
)

type Paymenter interface {
	Pay(amount string)
}

type CardPayment struct {
}

func (c CardPayment) Pay(amount string) {

	fmt.Println("the amount has been paid By Card: ", amount)
}

type CashPayment struct {
}

func (Cs CashPayment) Pay(amount string) {

	fmt.Println("the amount has been paid By Cash: ", amount)
}

type UPIPayment struct {
}

func (u UPIPayment) Pay(amount string) {

	fmt.Println("the amount has been paid By UPI: ", amount)
}

type Stripe struct {
}

func (S Stripe) Pay(amount string) {

	fmt.Println("the amount has been paid By Stripe: ", amount)
}

func main() {
	payments := []Paymenter{
		CashPayment{},
		CardPayment{},
		UPIPayment{},
		Stripe{},
	}

	for _, payment := range payments {
		payment.Pay("500")
	}
}
