package main

type Observer interface {
	notify(info any)
	getId() string
	getName() string
	getAge() int
}
