package main

import (
	pm "Payments/paymentMethods"
	"fmt"
)

func MakePayment(p pm.PaymentMethod, amt float64) (float64, error) {
	if err := p.Pay(amt); err != nil {
		fmt.Println("Payment failed:", err)
	}
	switch v := p.(type) {
	case *pm.CreditCard:
		return v.Limit, nil
	case *pm.DebitCard:
		return v.Balance, nil
	case *pm.UPI:
		return v.Balance, nil
	default:
		return 0, fmt.Errorf("unknown payment error")
	}
}

func main() {

	cc := &pm.CreditCard{Card: pm.Card{CardNumber: "1234", Name: "Akshat", CVV: "909"},
		Limit: 1000}
	dc := &pm.DebitCard{Card: pm.Card{CardNumber: "9080", Name: "AkshatDC", CVV: "323"},
		Balance: 2000}
	limit, err1 := MakePayment(cc, 200)
	if err1 != nil {
		fmt.Errorf("error in credit card payment")
	} else {
		fmt.Printf("Remaining limit in credit card %s: %.2f\n", cc.CardNumber, limit)
	}
	bal, err2 := MakePayment(dc, 1000)
	if err2 != nil {
		fmt.Errorf("error in debit card payment")
	} else {
		fmt.Printf("Remaining balance in debit card %s: %.2f\n", cc.CardNumber, bal)
	}

}
