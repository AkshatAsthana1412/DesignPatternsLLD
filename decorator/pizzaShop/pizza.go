package main

type Pizza interface {
	GetPrice() float32
}

// Concrete object
type BasicPizza struct {
	price float32
}

func (p *BasicPizza) GetPrice() float32 {
	return p.price
}
