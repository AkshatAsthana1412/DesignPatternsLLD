package paymentMethods

import (
	"fmt"
)

type UPI struct {
	UPIId   string
	Balance float64
}

func (upi *UPI) Pay(amt float64) error {
	fmt.Printf("Spent Rs. %.2f by UPI %s", amt, upi.UPIId)
	upi.Balance -= amt
	return nil
}
