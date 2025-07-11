package functions

import (
	"fmt"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/common"
	"go.uber.org/zap"
	"nextui-led-control/models"
	"os"
	"strings"
)

var IsBrick bool

func init() {
	IsBrick = checkIfBrick()
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

	chmod(filePath, 1)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		fmt.Fprintf(file, "%d\n", led.InfoBrightness)
		file.Close()
	}

	chmod(filePath, 0)

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

	chmod(filePath, 1)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		fmt.Fprintf(file, "%d\n", led.Brightness)
		file.Close()
	}

	chmod(filePath, 0)

	SetEffect(led)
}

func SetEffect(led models.LED) {
	filePath := fmt.Sprintf("/sys/class/led_anim/effect_%s", led.InternalName)

	chmod(filePath, 1)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		fmt.Fprintf(file, "%d\n", led.Effect)
		file.Close()
	}

	chmod(filePath, 0)
}

func SetEffectCycles(led models.LED) {
	filePath := fmt.Sprintf("/sys/class/led_anim/effect_cycles_%s", led.InternalName)

	chmod(filePath, 1)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		fmt.Fprintf(file, "%d\n", led.Trigger) // Using Trigger field for cycles
		file.Close()
	}

	chmod(filePath, 0)

	SetEffect(led)
}

func SetEffectSpeed(led models.LED) {
	filePath := fmt.Sprintf("/sys/class/led_anim/effect_duration_%s", led.InternalName)

	chmod(filePath, 1)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		fmt.Fprintf(file, "%d\n", led.Speed)
		file.Close()
	}

	chmod(filePath, 0)

	SetEffect(led)
}

func SetColor(led models.LED) {
	logger := common.GetLoggerInstance()
	filePath := fmt.Sprintf("/sys/class/led_anim/effect_rgb_hex_%s", led.InternalName)

	chmod(filePath, 1)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		logger.Error("Unable to save LED color", zap.Error(err))
	}

	if err == nil {
		fmt.Fprintf(file, "%s\n", led.Color1)
		file.Close()
	}

	chmod(filePath, 0)
	SetEffect(led)
}

func chmod(file string, writable int) {
	fileInfo, err := os.Stat(file)
	if err == nil {
		currentMode := fileInfo.Mode()
		var newMode os.FileMode

		if writable != 0 {
			newMode = currentMode | 0222 // S_IWUSR | S_IWGRP | S_IWOTH
		} else {
			newMode = currentMode &^ 0222 // Remove S_IWUSR | S_IWGRP | S_IWOTH
		}

		err := os.Chmod(file, newMode)
		if err != nil {
			fmt.Printf("chmod error %d %s\n", writable, file)
		}
	} else {
		fmt.Printf("stat error %d %s\n", writable, file)
	}
}
