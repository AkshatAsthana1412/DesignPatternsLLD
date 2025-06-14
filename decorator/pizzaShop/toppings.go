package main

type TomatoTopping struct {
	p Pizza
}

func (tt *TomatoTopping) GetPrice() float32 {
	return tt.p.GetPrice() + 5
}

type ExtraCheese struct {
	p Pizza
}

func (ct *ExtraCheese) GetPrice() float32 {
	return ct.p.GetPrice() + 10
}

type JalapenoTopping struct {
	p Pizza
}

func (jt *JalapenoTopping) GetPrice() float32 {
	return jt.p.GetPrice() + 8
}
