package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type Person interface {
	SayHello()
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, age %d\n", p.Name, p.Age)
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{Name: name, Age: age}
	} else {
		return &person{Name: name, Age: age}
	}

}

// with this design, if we want to add another type of person, we can just create another struct defining it
type tiredPerson struct {
	Name string
	Age  int
}

func (tp *tiredPerson) SayHello() {
	fmt.Println("I'm a tired person")
}

func main() {
	p := NewPerson("Akshat", 28) //name and age are encapsulated in person and cannot be modified
	// because p doesn't expose the fields of person
	q := NewPerson("John", 190)
	p.SayHello()
	q.SayHello()
}
