package syrops

// Inerface for syrops
// GetMl - gets the amount of syrop
// GetType - gets the syrop type
// AddMore - ads more syrop
type Syrop interface {
	// gets the amount of syrop
	GetMl() float64
	// gets the syrop type
	GetType() string
	// ads more syrop
	AddMore(float64) Syrop
}
