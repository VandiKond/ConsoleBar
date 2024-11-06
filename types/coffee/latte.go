package coffee

import (
	"fmt"

	Errors "github.com/VandiKond/ConsoleBar/errors"
	coffeerrors "github.com/VandiKond/ConsoleBar/errors/coffe_errors"
	"github.com/VandiKond/ConsoleBar/types/additives"
	"github.com/VandiKond/ConsoleBar/types/syrops"
)

// A coffe ctructure
// MlCoffe - the amount of black coffe
// MlMilk - the amount of milk
// Syrop - the syrop that's used
// Additives - a slice of additives
type Latte struct {
	// the amount of black coffe
	MlCoffe float64
	// the amount of milk
	MlMilk float64
	// the syrop that's used
	Syrop syrops.Syrop
	// a slice of additives
	Additives []additives.Additive
}

// Ads syrop to latte
//
// Returns an error if syrops don't match
func (l Latte) AddSyrop(syrop syrops.Syrop) (Coffee, error) {
	// Checks are syrops similar
	if l.Syrop.GetType() != syrop.GetType() {
		return l, Errors.NewError(coffeerrors.TDS, fmt.Sprintf("%s is not %s", l.Syrop.GetType(), syrop.GetType()))
	}
	// Ads syrop
	l.Syrop = l.Syrop.AddMore(syrop.GetMl())

	// returns the latte
	return l, nil
}

// Ads an additive
func (l Latte) AddAdditive(additive additives.Additive) (Coffee, error) {
	// Ads a new element to the slice of edditives
	l.Additives = append(l.Additives, additive)

	// returns the latte
	return l, nil
}
