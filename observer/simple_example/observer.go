package main

type Observer interface {
	update(info any)
	getId() string
}
