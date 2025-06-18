package ui

import (
	"fmt"
	"github.com/UncleJunVIP/gabagool/pkg/gabagool"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/common"
	"github.com/veandco/go-sdl2/sdl"
	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"nextui-led-control/functions"
	"nextui-led-control/models"
	"qlova.tech/sum"
	"strconv"
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
		fmt.Println("New effect: ", newValue)
		// Update LED effect and apply changes
		if !functions.IsDev() {
			effect := newValue.(models.EffectType)
			m.LED.Effect = int(effect)
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
						logger.Debug("New Color", zap.String("hexCode", hexCode))

						if !functions.IsDev() {
							logger.Debug("Setting Color", zap.String("hexCode", hexCode))
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

	items = append(items, gabagool.ItemWithOptions{
		Item: gabagool.MenuItem{
			Text: "Info Brightness",
		},
		Options:        infoBrightnessOptions,
		SelectedOption: infoBrightnessIndex,
	})

	footerItems := []gabagool.FooterHelpItem{
		{ButtonName: "B", HelpText: "Back"},
		{ButtonName: "Start", HelpText: "Save"},
	}

	title := cases.Title(language.English).String(m.LED.DisplayName)

	result, err := gabagool.OptionsList(fmt.Sprintf("%s Settings", title), items, footerItems)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !result.IsSome() || result.Unwrap().Canceled {
		return
	}

	selections := result.Unwrap()

	return selections, 0, nil
}
