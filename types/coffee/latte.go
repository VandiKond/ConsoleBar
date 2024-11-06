package coffee

import (
	"fmt"

	Errors "github.com/VandiKond/ConsoleBar/errors"
	coffeerrors "github.com/VandiKond/ConsoleBar/errors/coffe_errors"
	"github.com/VandiKond/ConsoleBar/types/additives"
	"github.com/VandiKond/ConsoleBar/types/syrups"
)

// A coffe ctructure
// MlCoffe - the amount of black coffe
// MlMilk - the amount of milk
// Syrup - the syrup that's used
// Additives - a slice of additives
type Latte struct {
	// the amount of black coffe
	MlCoffe float64
	// the amount of milk
	MlMilk float64
	// the syrup that's used
	Syrup syrups.Syrup
	// a slice of additives
	Additives []additives.Additive
}

// Adds syrup to latte
//
// Returns an error if syrups don't match
func (l Latte) AddSyrup(syrup syrups.Syrup) (Coffee, error) {
	// Checks are syrups similar
	if l.Syrup.GetType() != syrup.GetType() {
		return l, Errors.NewError(coffeerrors.TDS, fmt.Sprintf("%s is not %s", l.Syrup.GetType(), syrup.GetType()))
	}
	// Adds syrup
	l.Syrup = l.Syrup.AddMore(syrup.GetMl())

	// returns the latte
	return l, nil
}

// Adds an additive
func (l Latte) AddAdditive(additive additives.Additive) (Coffee, error) {
	// Adds a new element to the slice of additives
	l.Additives = append(l.Additives, additive)

	// returns the latte
	return l, nil
}
