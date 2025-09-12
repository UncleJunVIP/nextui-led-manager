package main

import (
	"fmt"
	"nextui-led-control/functions"
	"nextui-led-control/models"
	"nextui-led-control/ui"
	"os"

	gaba "github.com/UncleJunVIP/gabagool/pkg/gabagool"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/common"
)

var settings map[string]models.LED

func init() {
	common.SetLogLevel("DEBUG")
	gaba.InitSDL(gaba.GabagoolOptions{
		WindowTitle:    "LED Manager",
		ShowBackground: true,
	})

	logger := common.GetLoggerInstance()

	logger.Debug("LED Manager started")

	var err error

	settings, err = functions.LoadLEDSettings()

	if err != nil {
		fmt.Println("Error loading LED settings:", err)
		os.Exit(1)
	}

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
				sel := res.(models.LED)
				led := settings[sel.InternalName]
				screen = ui.InitLedSettings(led)
			case -1, 2:
				os.Exit(0)
			}
		case models.ScreenNames.LedSettings:
			switch code {
			case 0:
				led := res.(models.LED)
				settings[led.InternalName] = led

				var leds []models.LED
				for _, v := range settings {
					leds = append(leds, v)
				}

				functions.WriteLEDSettings(leds)
			}

			screen = ui.InitMainMenu()
		}
	}

}
