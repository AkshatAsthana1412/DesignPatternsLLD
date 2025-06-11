package main

import "fmt"

type Employee struct {
	name, position string
	salary         int
}

// Struct way of creating a factory
type EmployeeFactory struct {
	position string
	salary   int
}

func NewEmployeeFactory(position string, salary int) *EmployeeFactory {
	return &EmployeeFactory{position, salary}
}

func (ef *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name: name, position: ef.position, salary: ef.salary}
}

// Functional way of creating a factory
func NewEmployeeFactory2(position string, salary int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name: name, position: position, salary: salary}
	}
}

func main() {
	devFactory := NewEmployeeFactory("Developer", 80000)
	managerFactory := NewEmployeeFactory("Manager", 90000)
	hrFactory := NewEmployeeFactory2("HR", 30000)
	a := devFactory.Create("Akshat")
	b := managerFactory.Create("Alice")
	c := hrFactory("Helen")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
