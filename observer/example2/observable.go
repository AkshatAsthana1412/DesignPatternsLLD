package main

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	NotifyAllObservers()
}
