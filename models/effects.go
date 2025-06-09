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
	Rainbow
	Twinkle
	Fire
	Glitter
	NeonGlow
	Firefly
	Aurora
	Reactive
)

type TopbarEffectType int

const (
	TopbarLinear TopbarEffectType = iota
	TopbarBreathe
	TopbarIntervalBreathe
	TopbarStatic
	TopbarBlink1
	TopbarBlink2
	TopbarBlink3
	TopbarRainbow
	TopbarTwinkle
	TopbarFire
	TopbarGlitter
	TopbarNeonGlow
	TopbarFirefly
	TopbarAurora
	TopbarReactive
	TopbarRainbowSpecial
	TopbarNight
)

type LREffectType int

const (
	LRLinear LREffectType = iota
	LRBreathe
	LRIntervalBreathe
	LRStatic
	LRBlink1
	LRBlink2
	LRBlink3
	LRRainbow
	LRTwinkle
	LRFire
	LRGlitter
	LRNeonGlow
	LRFirefly
	LRAurora
	LRReactive
	LRRainbowSpecial
	LRReactiveSpecial
)

func (e EffectType) String() string {
	return [...]string{
		"Linear", "Breathe", "Interval Breathe", "Static",
		"Blink 1", "Blink 2", "Blink 3", "Rainbow", "Twinkle",
		"Fire", "Glitter", "NeonGlow", "Firefly", "Aurora", "Reactive",
	}[e]
}

func (e TopbarEffectType) String() string {
	return [...]string{
		"Linear", "Breathe", "Interval Breathe", "Static",
		"Blink 1", "Blink 2", "Blink 3", "Rainbow", "Twinkle",
		"Fire", "Glitter", "NeonGlow", "Firefly", "Aurora", "Reactive",
		"Topbar Rainbow", "Topbar night",
	}[e]
}

func (e LREffectType) String() string {
	return [...]string{
		"Linear", "Breathe", "Interval Breathe", "Static",
		"Blink 1", "Blink 2", "Blink 3", "Rainbow", "Twinkle",
		"Fire", "Glitter", "NeonGlow", "Firefly", "Aurora", "Reactive",
		"LR Rainbow", "LR Reactive",
	}[e]
}

func GetStandardEffectOptions() []gabagool.Option {
	options := make([]gabagool.Option, 15)
	for i := EffectType(0); i <= Reactive; i++ {
		options[i] = gabagool.Option{
			DisplayName: i.String(),
			Value:       i,
		}
	}
	return options
}

func GetTopbarEffectOptions() []gabagool.Option {
	options := make([]gabagool.Option, 17)
	for i := TopbarEffectType(0); i <= TopbarNight; i++ {
		options[i] = gabagool.Option{
			DisplayName: i.String(),
			Value:       i,
		}
	}
	return options
}

func GetLREffectOptions() []gabagool.Option {
	options := make([]gabagool.Option, 17)
	for i := LREffectType(0); i <= LRReactiveSpecial; i++ {
		options[i] = gabagool.Option{
			DisplayName: i.String(),
			Value:       i,
		}
	}
	return options
}
