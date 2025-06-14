package main

import "fmt"

func main() {
	pizza := &BasicPizza{price: 10}
	cheeseToppedPizza := &ExtraCheese{p: pizza}
	// without changing the BasicPizza struct, we're able to add toppings(behaviors) to it through decorators.
	jalapenoCheesePizza := JalapenoTopping{p: cheeseToppedPizza}
	finalPrice := jalapenoCheesePizza.GetPrice()
	fmt.Println("Price of Jalapeno Cheese Pizza: ", finalPrice)
}
