package main

import (
	"fmt"
	"slices"
)

type stockTicker struct {
	observers  []Observer
	tickerName string
	price      float64
}

func NewTicker(tn string, tp float64) *stockTicker {
	return &stockTicker{tickerName: tn, price: tp}
}

func (st *stockTicker) Register(o ...Observer) {
	st.observers = append(st.observers, o...)
}

func (st *stockTicker) Deregister(on ...string) {
	for _, obsName := range on {
		st.observers = slices.DeleteFunc(st.observers, func(obs Observer) bool {
			return obs.getName() == obsName
		})
	}
}

func (st *stockTicker) UpdatePrice(newPrice float64) {
	st.price = newPrice
	for _, obs := range st.observers {
		err := obs.notify(st.tickerName, newPrice)
		if err != nil {
			panic(fmt.Sprintf("Could not notify to all observers %s", obs.getName()))
		}
	}
}

func (st *stockTicker) PrintObservers() {
	for i, o := range st.observers {
		fmt.Printf("Observer %d: %s\n", i+1, o.getName())
	}
}
