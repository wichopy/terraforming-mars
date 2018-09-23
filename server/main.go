package main

import (
	// Need to name import beause repo has a dash
	"fmt"

	common "./common"
)

type App struct {
	Production      common.Production
	Supply          common.Supply
	TerraformRating uint8
}

func main() {
	app := App{
		// Define this by corporation cards.
		Production: common.Production{
			MegaCredits: 0,
			Steel:       0,
			Titanium:    0,
			Plants:      0,
			Energy:      0,
			Heat:        0,
		},
		Supply: common.Supply{
			MegaCredits: 0,
			Steel:       0,
			Titanium:    0,
			Plants:      0,
			Energy:      0,
			Heat:        0,
		},
		// Depends on game mode.
		TerraformRating: 20,
	}

	fmt.Printf("%+v\n", app)
}
