package main

import (
	"container/list"
	"fmt"
)

type Traffic_Management struct {
	persons *list.List
}

func (t *Traffic_Management) Subscribe(o ...Observer) {
	for _, observer := range o {
		fmt.Printf("Observing: %s\n", observer.getName())
		t.persons.PushBack(observer)
	}
}

func (t *Traffic_Management) Unsubscribe(o ...Observer) {
	for _, observer := range o {
		for ele := t.persons.Front(); ele != nil; ele = ele.Next() {
			if ele.Value.(Observer) == observer {
				t.persons.Remove(ele)
			}
		}
	}
}

func (t *Traffic_Management) NotifyAllObservers() {
	for observer := t.persons.Front(); observer != nil; observer = observer.Next() {
		if p, ok := observer.Value.(*Person); ok {
			if p.eligibilityCriteria.IsEligible(p) {
				p.notify(fmt.Sprintf("Congrats %s! you're allowed to drive now.", p.getName()))
			} else {
				p.notify(fmt.Sprintf("%s you're not yet eligible to drive", p.getName()))
			}
		} else {
			fmt.Println("Invalid observer type")
		}
	}
}
