package coffee

import (
	"github.com/VandiKond/ConsoleBar/types/additives"
	"github.com/VandiKond/ConsoleBar/types/syrups"
)

// An intarface for coffee
// AddSyrup - adds more syrup (returns the coffe with more syrup and an error)
// AddAdditive - adds one more additive (returns the coffe with more additives and an error)
type Coffee interface {
	// adds more syrup (returns the coffe with more syrup and an error)
	AddSyrup(syrups.Syrup) (Coffee, error)
	// AddAdditive - adds one more additive (returns the coffe with more additives and an error)
	AddAdditive(additives.Additive) (Coffee, error)
}
