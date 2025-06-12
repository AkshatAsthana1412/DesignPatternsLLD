package main

import (
	"container/list"
	"fmt"
)

type Item struct {
	itemName     string
	observerList *list.List
	inStock      bool
}

func NewItem(name string, inStock bool) *Item {
	return &Item{itemName: name, observerList: new(list.List), inStock: inStock}
}

func (i *Item) register(o ...Observer) {
	for _, obs := range o {
		fmt.Printf("%s subscribed by %s\n", i.itemName, obs.getId())
		i.observerList.PushBack(obs)
	}
}

func (i *Item) deregister(o ...Observer) {
	for _, observer := range o {
		for obs := i.observerList.Front(); obs != nil; obs = obs.Next() {
			if obs.Value.(Observer) == observer {
				fmt.Printf("%s unsubscribed from %s\n", observer.getId(), i.itemName)
				i.observerList.Remove(obs)
			}
		}
	}
}

func (i *Item) notifyAll() {
	fmt.Println("Notifying all users")
	for obs := i.observerList.Front(); obs != nil; obs = obs.Next() {
		obs.Value.(Observer).update(fmt.Sprintf("Hey %s! %s is in stock", obs.Value.(Observer).getId(), i.itemName))
	}
}

func (i *Item) updateAvailability() {
	i.inStock = true
	i.notifyAll()
}
