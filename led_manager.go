package main

import (
	"fmt"
	gaba "github.com/UncleJunVIP/gabagool/pkg/gabagool"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/common"
	"nextui-led-control/models"
	"nextui-led-control/ui"
	"os"
)

func init() {
	gaba.InitSDL(gaba.GabagoolOptions{
		WindowTitle:    "LED Manager",
		ShowBackground: true,
	})

	common.SetLogLevel("ERROR")
}

func cleanup() {
	common.CloseLogger()
	gaba.CloseSDL()
}

func main() {
	defer cleanup()

	var screen models.Screen

	screen = ui.InitMainMenu()

	for {
		res, code, _ := screen.Draw()

		switch screen.Name() {
		case models.ScreenNames.MainMenu:
			switch code {
			case 0:
				// Set LED Setting Screen Here
				fmt.Println(res)
				screen = ui.InitLedSettings(models.LED{})
			case -1, 2:
				os.Exit(0)
			}
		}
	}

}
