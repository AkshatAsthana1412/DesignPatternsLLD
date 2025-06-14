package main

// decorator decorating p Pizza
type TomatoTopping struct {
	p Pizza
}

func (tt *TomatoTopping) GetPrice() float32 {
	return tt.p.GetPrice() + 5
}

// another decorator
type ExtraCheese struct {
	p Pizza
}

func (ct *ExtraCheese) GetPrice() float32 {
	return ct.p.GetPrice() + 10
}

// decorator
type JalapenoTopping struct {
	p Pizza
}

func (jt *JalapenoTopping) GetPrice() float32 {
	return jt.p.GetPrice() + 8
}
