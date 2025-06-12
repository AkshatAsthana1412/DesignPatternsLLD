package main

func main() {
	c1 := NewCustomer("Shayan", "1233")
	c2 := NewCustomer("Sharan", "98932")
	item := NewItem("Shirt", false)
	item.register(c1, c2)
	item.updateAvailability()
}
