package additives

// A inerfase for additives
// GetGrams - sends the gramm amount
// AddMore - ads more grams of the additive
// GetType - sends the type of the additive
type Additive interface {
	// sends the gramm amount
	GetGrams() float64
	// ads more grams of the additive
	AddMore(float64) Additive
	// sends the type of the additive
	GetType() string
}