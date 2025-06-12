package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Person struct {
	age                 int
	Name                string
	id                  string
	eligibilityCriteria Eligibility
}

func getRandomId() int {
	rand.Seed(time.Now().UnixNano()) // Seed to ensure different results on each run
	return rand.Intn(91) + 10        // 91 = 100 - 10 + 1
}

func NewPerson(name string, age int) *Person {
	return &Person{age, name, strconv.Itoa(getRandomId()), &AgeEligibility{}}
}

// to notify the person that he/she is allowed to drive now
func (p *Person) notify(info any) {
	fmt.Println(info.(string))
}

func (p *Person) getId() string {
	return p.id
}

func (p *Person) getName() string {
	return p.Name
}

func (p *Person) getAge() int {
	return p.age
}

func (p *Person) setEligibility(e Eligibility) {
	p.eligibilityCriteria = e
}

func (p *Person) setAge(newAge int) {
	fmt.Printf("%s new age set to: %d\n", p.getName(), newAge)
	p.age = newAge
}
