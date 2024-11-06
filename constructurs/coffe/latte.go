package coffe_construcurs

import (
	"fmt"
	"strconv"

	Errors "github.com/VandiKond/ConsoleBar/errors"
	coffeerrors "github.com/VandiKond/ConsoleBar/errors/coffe_errors"
	syroperros "github.com/VandiKond/ConsoleBar/errors/syrop_erros"
	"github.com/VandiKond/ConsoleBar/types/additives"
	"github.com/VandiKond/ConsoleBar/types/coffee"
	"github.com/VandiKond/ConsoleBar/types/syrops"
)

// # Creating a new latte in the console
// Getting ml of balck coffe
// ml of milk
// The type and amount (in ml) of syrop
// Creating an array of additives for the coffee
// Sending the informaition about the coffe creation and returning a created latte
func NewLatte() (coffee.Latte, error) {
	// Starting the creation
	fmt.Println("Hi! Let's Make your latte!\nIf you want to stop say 'exit'")

	// Askong for exit
	var exit string
	fmt.Scanln(&exit)
	if exit == "exit" {
		return coffee.Latte{}, fmt.Errorf("user exited")
	}

	// Getting and validating the balck coffe amount
	fmt.Println("Chose the amount of black coffe you want")
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

	// getting the syrop type
	fmt.Println("Write a syrop type here. Chose between:\nvanila\n\nafter say the amount of syrop")
	var syropType string
	var syropAmount float64
	var Syrop syrops.Syrop
	fmt.Scanln(&syropType)

	// validating the syrop type
	switch syropType {
	case "vanila":
		Syrop = &syrops.Vanila{}
		break
	default:
		fmt.Println("Unknown syrop type: ", syropType)
		return coffee.Latte{}, Errors.NewError(syroperros.UST, syropType)
	}

	fmt.Println("Now write the amount of syrop you want (in ml)")
	fmt.Scanln(&syropAmount)
	Syrop = Syrop.AddMore(syropAmount)

	// Creating the base latte without the additives
	latte := coffee.Latte{MlCoffe: coffeMl, MlMilk: milkMl, Syrop: Syrop}

	// Asking for the attigives and going in the loop
	additiveSl := []additives.Additive{}
	fmt.Println("Now, please write the additives you want:\n\nsugar\nafter each additive write the amount of it\nfor exiting write 'stop'")
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
			fmt.Println("Unknown additive: ", addtiveType, "\nContinuing, if you wnat to stop write 'stop'")
			continue
		}
		if breakQ {
			break
		}

		// getting the additive amount
		var additiveAmmount float64
		fmt.Println("now write the ammount of", addtiveType, "you want (in gramms)")
		fmt.Scanln(&additiveAmmount)

		// Adding the additive to the slice
		fmt.Println("Continue, if you want to stop adding additives write 'stop'")
		additive = additive.AddMore(additiveAmmount)
		additiveSl = append(additiveSl, additive)
	}
	// Adding all additives
	latte.Additives = additiveSl

	// Sending the coffee info
	fmt.Printf("Yey! You've made your latte:\nml of black coffe : %v\nml of milk : %v\nsyrop : %s (%v ml)\nadditives : \n", latte.MlCoffe, latte.MlMilk, latte.Syrop.GetType(), latte.Syrop.GetMl())
	for _, a := range latte.Additives {
		fmt.Printf("%s (%v gramms)\n", a.GetType(), a.GetGrams())
	}

	// Returning
	return latte, nil
}
