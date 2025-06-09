package main

import (
	"fmt"
)

// Example of aggregation (has-a relationship). Professor exists within university but can also
// exist independently.

type Professor struct {
	Name    string
	Subject string
}

func (p Professor) Teach() {
	fmt.Printf("%s is teaching %s\n", p.Name, p.Subject)
}

type University struct {
	Name       string
	Professors []*Professor
}

func (u *University) AddProfessor(p *Professor) {
	u.Professors = append(u.Professors, p)
}

func (u University) ShowProfessors() {
	fmt.Printf("Professors at %s\n", u.Name)
	for _, professor := range u.Professors {
		fmt.Printf(" - %s\n", professor.Name)
	}
}

func main() {
	prof1 := Professor{Name: "Akshat", Subject: "Quant Finance"}
	prof2 := Professor{Name: "Alan", Subject: "Math"}
	uni := University{Name: "MIT"}
	uni.AddProfessor(&prof1)
	uni.AddProfessor(&prof2)
	uni.ShowProfessors()
	prof1.Teach()
}
