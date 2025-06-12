package main

import "container/list"

func main() {
	p1 := NewPerson("Akshat", 28)
	p2 := NewPerson("Yashika", 17)
	tf := Traffic_Management{new(list.List)}
	tf.Subscribe(p1, p2)
	tf.NotifyAllObservers()
	p2.setAge(20)
	tf.NotifyAllObservers()
}
