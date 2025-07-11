package models

import (
	"github.com/UncleJunVIP/gabagool/pkg/gabagool"
)

var EffectNames = []string{
	"Linear", "Breathe", "Interval Breathe", "Static",
	"Blink 1", "Blink 2",
}

func GetStandardEffectOptions(onUpdate func(newValue interface{})) []gabagool.Option {
	var options []gabagool.Option

	for i, name := range EffectNames {
		options = append(options, gabagool.Option{
			DisplayName: name,
			Value:       i + 1,
			OnUpdate:    onUpdate,
		})
	}
	return options
}
