package syrups

// A syrup
// Ml - the amount of the syrup
type Vanila struct {
	// the amount of the syrup
	Ml float64
}

// gets the amount of vanila syrup
func (v Vanila) GetMl() float64 {
	return v.Ml
}

// sends the type of the vanila syrup
func (v Vanila) GetType() string {
	return "vanila"
}

// adds more vanila syrup
func (v Vanila) AddMore(ml float64) Syrup {
	v.Ml += ml
	return v
}
