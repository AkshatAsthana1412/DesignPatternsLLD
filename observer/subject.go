package main

type Subject interface {
	register(o Observer)
	deregister(o Observer)
	notifyAll()
}
