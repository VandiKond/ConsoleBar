package coffe_construcurs

import (
	"fmt"
	"strconv"

	Errors "github.com/VandiKond/ConsoleBar/errors"
	coffeerrors "github.com/VandiKond/ConsoleBar/errors/coffe_errors"
	syruperros "github.com/VandiKond/ConsoleBar/errors/syrup_erros"
	"github.com/VandiKond/ConsoleBar/types/additives"
	"github.com/VandiKond/ConsoleBar/types/coffee"
	"github.com/VandiKond/ConsoleBar/types/syrups"
)

// # Creating a new latte in the console
// Getting ml of black coffee
// ml of milk
// The type and amount (in ml) of syrup
// Creating an array of additives for the coffee
// Sending the information about the coffee creation and returning a created latte
func NewLatte() (coffee.Latte, error) {
	// Starting the creation
	fmt.Println("Hi! Let's Make your latte!\nIf you want to stop say 'exit'")

	// Asking for exit
	var exit string
	fmt.Scanln(&exit)
	if exit == "exit" {
		return coffee.Latte{}, fmt.Errorf("user exited")
	}

	// Getting and validating the black coffee amount
	fmt.Println("Chose the amount of black coffee you want")
	var coffeMl float64
	var baseFloat float64
	fmt.Scanln(&coffeMl)
	if coffeMl == baseFloat || coffeMl <= 0 {
		fmt.Println("Please write a valid amount")
		return coffee.Latte{}, Errors.NewError(coffeerrors.NEC, strconv.FormatFloat(coffeMl, 'f', -1, 64))
	}

	// Getting the milk amount
	fmt.Println("Chose the amount of milk you want")
	var milkMl float64
	fmt.Scanln(&milkMl)

	// getting the syrup type
	fmt.Println("Write a syrup type here. Chose between:\nvanilla\n\nafter say the amount of syrup")
	var syrupType string
	var syrupAmount float64
	var Syrup syrups.Syrup
	fmt.Scanln(&syrupType)

	// validating the syrup type
	switch syrupType {
	case "vanilla":
		Syrup = &syrups.Vanilla{}
		break
	default:
		fmt.Println("Unknown syrup type: ", syrupType)
		return coffee.Latte{}, Errors.NewError(syruperros.UST, syrupType)
	}

	fmt.Println("Now write the amount of syrup you want (in ml)")
	fmt.Scanln(&syrupAmount)
	Syrup = Syrup.AddMore(syrupAmount)

	// Creating the base latte without the additives
	latte := coffee.Latte{MlCoffe: coffeMl, MlMilk: milkMl, Syrup: Syrup}

	// Asking for the additives and going in the loop
	additiveSl := []additives.Additive{}
	fmt.Println("Now, please write the additives you want:\n\nsugar \nafter each additive write the amount of it\nfor exiting write 'stop'")
	for {
		// Getting the additive
		breakQ := false
		var addtiveType string
		var additive additives.Additive
		fmt.Scanln(&addtiveType)

		// Validating the additive
		switch addtiveType {
		case "sugar":
			additive = &additives.Sugar{}
			break
		case "stop":
			breakQ = true
			break
		default:
			fmt.Println("Unknown additive: ", addtiveType, "\nContinuing, if you want to stop write 'stop'")
			continue
		}
		if breakQ {
			break
		}

		// getting the additive amount
		var additiveAmmount float64
		fmt.Println("now write the amount of", addtiveType, "you want (in grams)")
		fmt.Scanln(&additiveAmmount)

		// Adding the additive to the slice
		fmt.Println("Continue, if you want to stop adding additives write 'stop'")
		additive = additive.AddMore(additiveAmmount)
		additiveSl = append(additiveSl, additive)
	}
	// Adding all additives
	latte.Additives = additiveSl

	// Sending the coffee info
	fmt.Printf("Yey! You've made your latte:\nml of black coffee : %v\nml of milk : %v\nsyrup : %s (%v ml)\nadditives : \n", latte.MlCoffe, latte.MlMilk, latte.Syrup.GetType(), latte.Syrup.GetMl())
	for _, a := range latte.Additives {
		fmt.Printf("%s (%v grams)\n", a.GetType(), a.GetGrams())
	}

	// Returning
	return latte, nil
}
