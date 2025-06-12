package main

import "fmt"

type Customer struct {
	name string
	id   string
}

func (c *Customer) update(info any) {
	fmt.Println(info.(string))
}

func NewCustomer(name string, id string) *Customer {
	return &Customer{name, id}
}

func (c *Customer) getId() string {
	return c.id
}
