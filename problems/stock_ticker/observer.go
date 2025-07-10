package main

import "fmt"

type Observer interface {
	notify(ticker string, price float64) error
	getName() string
}

type LoggerObserver struct {
	name string
}

func (l *LoggerObserver) notify(tn string, tp float64) error {
	fmt.Printf("[LOGGER]Price of %s changed to %.2f\n", tn, tp)
	return nil
}

func (l *LoggerObserver) getName() string {
	return l.name
}

type AlertObserver struct {
	name     string
	strategy AlertStrategy
}

func (ao *AlertObserver) notify(tn string, tp float64) error {
	if ao.strategy.ShouldAlert(tn, tp) {
		fmt.Printf("Alert triggered for %s\n", tn)
	}
	return nil
}

func (ao *AlertObserver) getName() string {
	return ao.name
}
