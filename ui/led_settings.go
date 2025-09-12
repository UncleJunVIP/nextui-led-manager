package ui

import (
	"fmt"
	"nextui-led-control/functions"
	"nextui-led-control/models"
	"strconv"
	"strings"

	"github.com/UncleJunVIP/gabagool/pkg/gabagool"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/common"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func (m LedSettings) Draw() (settings interface{}, exitCode int, e error) {
	logger := common.GetLoggerInstance()

	logger.Debug("Drawing LED Settings")

	effectOptions := models.GetStandardEffectOptions(func(newValue interface{}) {
		fmt.Println(fmt.Sprintf("New effect: %d | %s", newValue, models.EffectNames[newValue.(int)-1]))
		if !functions.IsDev() {
			m.LED.Effect = newValue.(int)
			functions.SetEffect(m.LED)
		}
	})

	var colorValue sdl.Color
	if len(m.LED.Color1) > 0 {
		hexColor, err := strconv.ParseUint(m.LED.Color1, 16, 32)
		if err == nil {
			colorValue = sdl.Color{
				R: uint8((hexColor >> 16) & 0xFF),
				G: uint8((hexColor >> 8) & 0xFF),
				B: uint8(hexColor & 0xFF),
				A: 255,
			}
		} else {
			colorValue = sdl.Color{R: 255, G: 0, B: 0, A: 255}
		}
	} else {
		colorValue = sdl.Color{R: 255, G: 0, B: 0, A: 255}
	}

	items := []gabagool.ItemWithOptions{
		{
			Item: gabagool.MenuItem{
				Text:     "Color",
				Selected: true,
			},
			Options: []gabagool.Option{
				{
					DisplayName: fmt.Sprintf("#%s", m.LED.Color1),
					Value:       colorValue,
					Type:        gabagool.OptionTypeColorPicker,
					OnUpdate: func(newValue interface{}) {
						color := newValue.(sdl.Color)
						hexCode := fmt.Sprintf("%02X%02X%02X", color.R, color.G, color.B)
						logger.Debug("New Color", "hexCode", hexCode)

						if !functions.IsDev() {
							logger.Debug("Setting Color", "hexCode", hexCode)
							m.LED.Color1 = hexCode
							functions.SetColor(m.LED)
						}
					},
				},
			},
			SelectedOption: 0,
		},
		{
			Item: gabagool.MenuItem{
				Text: "Effect",
			},
			Options:        effectOptions,
			SelectedOption: m.LED.Effect - 1,
		},
	}

	speedOptions := make([]gabagool.Option, 51)
	for i := 0; i <= 50; i++ {
		value := i * 100
		speedOptions[i] = gabagool.Option{
			DisplayName: fmt.Sprintf("%d", value),
			Value:       value,
			Type:        gabagool.OptionTypeStandard,
		}
	}

	speedIndex := 10
	for i, opt := range speedOptions {
		if opt.Value.(int) == m.LED.Speed {
			speedIndex = i
			break
		}
	}

	items = append(items, gabagool.ItemWithOptions{
		Item: gabagool.MenuItem{
			Text: "Speed",
		},
		Options:        speedOptions,
		SelectedOption: speedIndex,
	})

	brightnessOptions := make([]gabagool.Option, 21)
	for i := 0; i <= 20; i++ {
		value := i * 5
		brightnessOptions[i] = gabagool.Option{
			DisplayName: fmt.Sprintf("%d%%", value),
			Value:       value,
			Type:        gabagool.OptionTypeStandard,
			OnUpdate: func(newValue interface{}) {
				m.LED.Brightness = newValue.(int)
				functions.SetBrightness(m.LED)
			},
		}
	}

	brightnessIndex := 10
	for i, opt := range brightnessOptions {
		if opt.Value.(int) == m.LED.Brightness {
			brightnessIndex = i
			break
		}
	}

	items = append(items, gabagool.ItemWithOptions{
		Item: gabagool.MenuItem{
			Text: "Brightness",
		},
		Options:        brightnessOptions,
		SelectedOption: brightnessIndex,
	})

	infoBrightnessOptions := make([]gabagool.Option, 21)
	for i := 0; i <= 20; i++ {
		value := i * 5
		infoBrightnessOptions[i] = gabagool.Option{
			DisplayName: fmt.Sprintf("%d%%", value),
			Value:       value,
			Type:        gabagool.OptionTypeStandard,
			OnUpdate: func(newValue interface{}) {
				m.LED.Brightness = newValue.(int)
				functions.SetInfoBrightness(m.LED)
			},
		}
	}

	// Info brightness index
	infoBrightnessIndex := 10 // Default value
	for i, opt := range brightnessOptions {
		if opt.Value.(int) == m.LED.InfoBrightness {
			infoBrightnessIndex = i
			break
		}
	}

	if !functions.IsKidMode() {
		items = append(items, gabagool.ItemWithOptions{
			Item: gabagool.MenuItem{
				Text: "Info Brightness",
			},
			Options:        infoBrightnessOptions,
			SelectedOption: infoBrightnessIndex,
		})
	}

	footerItems := []gabagool.FooterHelpItem{
		{ButtonName: "B", HelpText: "Back"},
		{ButtonName: "Start", HelpText: "Save"},
	}

	title := cases.Title(language.English).String(m.LED.DisplayName)

	result, err := gabagool.OptionsList(fmt.Sprintf("%s Settings", title), items, footerItems)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, 1, nil
	}

	if !result.IsSome() || result.Unwrap().Canceled {
		return nil, 1, nil
	}

	selections := result.Unwrap()

	newInfoBrightness := 50

	if !functions.IsKidMode() {
		newInfoBrightness = selections.Items[4].SelectedOption * 5
	}

	newSettings := models.LED{
		DisplayName:    m.LED.DisplayName,
		InternalName:   m.LED.InternalName,
		Color1:         strings.ReplaceAll(selections.Items[0].Options[0].DisplayName, "#", ""),
		Color2:         strings.ReplaceAll(selections.Items[0].Options[0].DisplayName, "#", ""),
		Effect:         selections.Items[1].SelectedOption + 1,
		Speed:          selections.Items[2].SelectedOption * 100,
		Brightness:     selections.Items[3].SelectedOption * 5,
		InfoBrightness: newInfoBrightness,
		Trigger:        1,
	}

	return newSettings, 0, nil
}
