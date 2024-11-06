package syrops

// A syrop
// Ml - the amount of the syrop
type Vanila struct {
	// the amount of the syrop
	Ml float64
}

// gets the amount of vanila syrop
func (v Vanila) GetMl() float64 {
	return v.Ml
}

// sends the type of the vanila syrop
func (v Vanila) GetType() string {
	return "vanila"
}

// ads more vanila syrop
func (v Vanila) AddMore(ml float64) Syrop {
	v.Ml += ml
	return v
}
