package main

import "container/list"

func main() {
	p1 := NewPerson("Akshat", 28)
	p2 := NewPerson("Yashika", 20)
	tf := Traffic_Management{new(list.List)}
	tf.Subscribe(p1, p2)
	tf.NotifyAllObservers()
}
