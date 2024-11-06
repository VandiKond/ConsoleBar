package coffee

import (
	"github.com/VandiKond/ConsoleBar/types/additives"
	"github.com/VandiKond/ConsoleBar/types/syrops"
)

// An intarface for coffee
// AddSyrop - ads more syrop (returns the coffe with more syrop and an error)
// AddAdditive - ads one more additive (returns the coffe with more additives and an error)
type Coffee interface {
	// ads more syrop (returns the coffe with more syrop and an error)
	AddSyrop(syrops.Syrop) (Coffee, error)
	// AddAdditive - ads one more additive (returns the coffe with more additives and an error)
	AddAdditive(additives.Additive) (Coffee, error)
}
