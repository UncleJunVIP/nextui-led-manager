package functions

import (
	"fmt"
	"nextui-led-control/models"
	"os"
	"path/filepath"
	"strings"

	"github.com/UncleJunVIP/nextui-pak-shared-functions/common"
)

var IsBrick bool

func init() {
	IsBrick = checkIfBrick()
	chmodLEDs()
}

func checkIfBrick() (containsBrick bool) {
	if os.Getenv("ENVIRONMENT") == "DEV" && os.Getenv("DEVICE") == "BRICK" {
		return true
	}

	filePath := "/usr/trimui/bin/MainUI"

	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return false
	}

	containsBrick = strings.Contains(string(content), "Brick")

	return containsBrick
}

func UpdateLED(led models.LED) {
	SetBrightness(led)
	SetEffect(led)
	SetColor(led)
	SetEffectCycles(led)
	SetEffectSpeed(led)
	SetInfoBrightness(led)
}

func SetInfoBrightness(led models.LED) {
	var filePath string

	if IsBrick {
		switch led.InternalName {
		case "m":
			filePath = "/sys/class/led_anim/max_scale"
		case "f1", "f2":
			filePath = "/sys/class/led_anim/max_scale_f1f2"
		default:
			filePath = fmt.Sprintf("/sys/class/led_anim/max_scale_%s", led.InternalName)
		}
	} else {
		filePath = "/sys/class/led_anim/max_scale"
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		fmt.Fprintf(file, "%d\n", led.InfoBrightness)
		file.Close()
	}

	SetEffect(led)
}

func SetBrightness(led models.LED) {
	var filePath string

	if IsBrick {
		switch led.InternalName {
		case "m":
			filePath = "/sys/class/led_anim/max_scale"
		case "f1", "f2":
			filePath = "/sys/class/led_anim/max_scale_f1f2"
		default:
			filePath = fmt.Sprintf("/sys/class/led_anim/max_scale_%s", led.InternalName)
		}
	} else {
		filePath = "/sys/class/led_anim/max_scale"
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		fmt.Fprintf(file, "%d\n", led.Brightness)
		file.Close()
	}

	SetEffect(led)
}

func SetEffect(led models.LED) {
	filePath := fmt.Sprintf("/sys/class/led_anim/effect_%s", led.InternalName)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		fmt.Fprintf(file, "%d\n", led.Effect)
		file.Close()
	}

}

func SetEffectCycles(led models.LED) {
	filePath := fmt.Sprintf("/sys/class/led_anim/effect_cycles_%s", led.InternalName)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		fmt.Fprintf(file, "%d\n", led.Trigger) // Using Trigger field for cycles
		file.Close()
	}

	SetEffect(led)
}

func SetEffectSpeed(led models.LED) {
	filePath := fmt.Sprintf("/sys/class/led_anim/effect_duration_%s", led.InternalName)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		fmt.Fprintf(file, "%d\n", led.Speed)
		file.Close()
	}

	SetEffect(led)
}

func SetColor(led models.LED) {
	logger := common.GetLoggerInstance()
	filePath := fmt.Sprintf("/sys/class/led_anim/effect_rgb_hex_%s", led.InternalName)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		logger.Error("Unable to save LED color", "error", err)
	}

	if err == nil {
		fmt.Fprintf(file, "%s\n", led.Color1)
		file.Close()
	}

	SetEffect(led)
}

func chmodLEDs() error {
	ledAnimPath := "/sys/class/led_anim"

	entries, err := os.ReadDir(ledAnimPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		filePath := filepath.Join(ledAnimPath, entry.Name())

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			continue
		}

		newMode := fileInfo.Mode() | 0222

		err = os.Chmod(filePath, newMode)
		if err != nil {
			fmt.Printf("chmod error on %s: %v\n", filePath, err)
		}
	}

	return nil
}
