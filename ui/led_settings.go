package ui

import (
	"fmt"
	"github.com/UncleJunVIP/gabagool/pkg/gabagool"
	"github.com/veandco/go-sdl2/sdl"
	"nextui-led-control/models"
	"qlova.tech/sum"
)

type LedSettings struct {
	LED models.LED
}

func InitLedSettings(led models.LED) LedSettings {
	return LedSettings{
		LED: led,
	}
}

func (m LedSettings) Name() sum.Int[models.ScreenName] {
	return models.ScreenNames.LedSettings
}

func (m LedSettings) Draw() (host interface{}, exitCode int, e error) {
	items := []gabagool.ItemWithOptions{
		{
			Item: gabagool.MenuItem{
				Text:     "Color",
				Selected: false,
			},
			Options: []gabagool.Option{
				{
					DisplayName: "#FF0000", // Default red color
					Value:       sdl.Color{R: 255, G: 0, B: 0, A: 255},
					Type:        gabagool.OptionTypeColorPicker,
				},
			},
			SelectedOption: 0,
		},
		{
			Item: gabagool.MenuItem{
				Text:     "Effect",
				Selected: true,
			},
			Options: models.GetStandardEffectOptions(),
		},
	}

	speedOptions := make([]gabagool.Option, 51) // 0, 100, 200, ..., 5000 (51 values)
	for i := 0; i <= 50; i++ {
		value := i * 100
		speedOptions[i] = gabagool.Option{
			DisplayName: fmt.Sprintf("%d", value),
			Value:       value,
			Type:        gabagool.OptionTypeStandard,
		}
	}

	items = append(items, gabagool.ItemWithOptions{
		Item: gabagool.MenuItem{
			Text:     "Speed",
			Selected: false,
		},
		Options:        speedOptions,
		SelectedOption: 10,
	})

	brightnessOptions := make([]gabagool.Option, 21)
	for i := 0; i <= 20; i++ {
		value := i * 5
		brightnessOptions[i] = gabagool.Option{
			DisplayName: fmt.Sprintf("%d%%", value),
			Value:       value,
			Type:        gabagool.OptionTypeStandard,
		}
	}

	items = append(items, gabagool.ItemWithOptions{
		Item: gabagool.MenuItem{
			Text:     "Brightness",
			Selected: false,
		},
		Options:        brightnessOptions,
		SelectedOption: 10,
	})

	items = append(items, gabagool.ItemWithOptions{
		Item: gabagool.MenuItem{
			Text:     "Info Brightness",
			Selected: false,
		},
		Options:        brightnessOptions,
		SelectedOption: 10,
	})

	footerItems := []gabagool.FooterHelpItem{
		{ButtonName: "B", HelpText: "Back"},
		{ButtonName: "Start", HelpText: "Save"},
	}

	result, err := gabagool.OptionsList("Top Bar Settings", items, footerItems)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Handle the result
	if !result.IsSome() || result.Unwrap().Canceled {
		fmt.Println("Selection canceled")
		return
	}

	// Get the selected colors
	selectedResult := result.Unwrap()
	for _, item := range selectedResult.Items {
		if len(item.Options) > 0 {
			if color, ok := item.Options[item.SelectedOption].Value.(sdl.Color); ok {
				fmt.Printf("%s: #%02X%02X%02X\n",
					item.Item.Text,
					color.R, color.G, color.B)
			}
		}
	}

	return nil, 0, nil
}
