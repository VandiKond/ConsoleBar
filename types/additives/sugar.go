package additives

// an additive
// Grams - amount of the additive
type Sugar struct {
	// amount of the additive
	Grams float64
}

// Gets the amount of grams
func (s Sugar) GetGrams() float64 {
	return s.Grams
}

// Sends the type of the additive
func (s Sugar) GetType() string {
	return "sugar"
}

// Ads more sugar
func (s Sugar) AddMore(grams float64) Additive {
	s.Grams += grams
	return s
}
