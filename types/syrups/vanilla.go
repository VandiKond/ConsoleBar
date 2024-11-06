package syrups

// A syrup
// Ml - the amount of the syrup
type Vanilla struct {
	// the amount of the syrup
	Ml float64
}

// gets the amount of vanilla syrup
func (v Vanilla) GetMl() float64 {
	return v.Ml
}

// sends the type of the vanilla syrup
func (v Vanilla) GetType() string {
	return "vanilla"
}

// adds more vanilla syrup
func (v Vanilla) AddMore(ml float64) Syrup {
	v.Ml += ml
	return v
}
