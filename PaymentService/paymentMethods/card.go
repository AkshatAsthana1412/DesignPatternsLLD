package paymentMethods

import (
	"fmt"
)

type Card struct {
	CardNumber string
	Name       string
	CVV        string
}

type CreditCard struct {
	Card
	Limit float64
}

func (cc *CreditCard) Pay(amt float64) error {
	fmt.Printf("Spent Rs. %.2f on credit card %s\n", amt, cc.CardNumber)
	cc.Limit -= amt
	return nil
}

type DebitCard struct {
	Card
	Balance float64
}

func (dc *DebitCard) Pay(amt float64) error {
	fmt.Printf("Spent Rs. %.2f on debit card %s\n", amt, dc.CardNumber)
	dc.Balance -= amt
	return nil
}
