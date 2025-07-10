package main

import "fmt"

func main() {
	sts := []*stockTicker{
		NewTicker("AAPL", 340),
		NewTicker("GOOG", 900),
		NewTicker("EFS", 200)}

	thresholds := map[string]float64{
		"AAPL": 350,
		"GOOG": 1000,
		"EFS":  100,
	}

	threshold_st := NewThresholdAlert(thresholds)
	lo := &LoggerObserver{name: "FirstLogger"}
	ao := &AlertObserver{name: "FirstAlert", strategy: threshold_st}

	for _, st := range sts {
		st.Register(lo, ao)
	}

	fmt.Printf("Subscribers of %s:\n", sts[0].tickerName)
	sts[0].PrintObservers()

	sts[0].UpdatePrice(399)
	sts[1].UpdatePrice(999)
}
