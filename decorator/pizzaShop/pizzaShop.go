package main

import "fmt"

func main() {
	pizza := &BasicPizza{price: 10}
	cheeseToppedPizza := &ExtraCheese{p: pizza}
	jalapenoCheesePizza := JalapenoTopping{p: cheeseToppedPizza}
	finalPrice := jalapenoCheesePizza.GetPrice()
	fmt.Println("Price of Jalapeno Cheese Pizza: ", finalPrice)
}
