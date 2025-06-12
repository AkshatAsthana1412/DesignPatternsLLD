package main

type Eligibility interface {
	IsEligible(o Observer) bool
}

type AgeEligibility struct {
}

func (ae *AgeEligibility) IsEligible(o Observer) bool {
	age := o.getAge()
	return age > 18
}
