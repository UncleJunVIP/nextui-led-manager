package models

import "github.com/UncleJunVIP/gabagool/pkg/gabagool"

type EffectType int

const (
	Linear EffectType = iota
	Breathe
	IntervalBreathe
	Static
	Blink1
	Blink2
	Blink3
)

func (e EffectType) String() string {
	return [...]string{
		"Linear", "Breathe", "Interval Breathe", "Static",
		"Blink 1", "Blink 2", "Blink 3",
	}[e]
}

func GetStandardEffectOptions(onUpdate func(newValue interface{})) []gabagool.Option {
	options := make([]gabagool.Option, 7)
	for i := EffectType(0); i <= Blink3; i++ {
		options[i] = gabagool.Option{
			DisplayName: i.String(),
			Value:       i + 1,
			OnUpdate:    onUpdate,
		}
	}
	return options
}
