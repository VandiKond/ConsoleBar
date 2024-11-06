package syrups

// Inerface for syrups
// GetMl - gets the amount of syrup
// GetType - gets the syrup type
// AddMore - adds more syrup
type Syrup interface {
	// gets the amount of syrup
	GetMl() float64
	// gets the syrup type
	GetType() string
	// adds more syrup
	AddMore(float64) Syrup
}
