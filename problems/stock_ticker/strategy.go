package main

type AlertStrategy interface {
	ShouldAlert(ticker string, price float64) bool
}

type ThresholdAlert struct {
	thresholds map[string]float64
}

func (ts *ThresholdAlert) ShouldAlert(ticker string, price float64) bool {
	if ts.thresholds[ticker] > price {
		return true
	} else {
		return false
	}
}

func NewThresholdAlert(ts map[string]float64) *ThresholdAlert {
	return &ThresholdAlert{thresholds: ts}
}
