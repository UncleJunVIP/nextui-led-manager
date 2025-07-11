package functions

import (
	"bufio"
	"nextui-led-control/models"
	"os"
	"strconv"
	"strings"
)

func IsDev() bool {
	return os.Getenv("ENVIRONMENT") == "DEV"
}

func GetSettingsFile() string {
	settingsFile := "/mnt/SDCARD/.userdata/shared/ledsettings.txt"
	if IsBrick {
		settingsFile = "/mnt/SDCARD/.userdata/shared/ledsettings_brick.txt"
	}

	if IsDev() {
		if IsBrick {
			settingsFile = "dev/ledsettings_brick.txt"
		} else {
			settingsFile = "dev/ledsettings.txt"
		}
	}

	return settingsFile
}

func LoadLEDSettings() (map[string]models.LED, error) {
	file, err := os.Open(GetSettingsFile())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var leds []models.LED
	scanner := bufio.NewScanner(file)

	var currentLED *models.LED
	var currentDisplayName string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines
		if line == "" {
			continue
		}

		// Check if this is a section header
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			// Save the previous LED if there was one
			if currentLED != nil {
				leds = append(leds, *currentLED)
			}

			// Start a new LED
			currentDisplayName = line[1 : len(line)-1] // Remove brackets
			currentLED = &models.LED{
				DisplayName: currentDisplayName,
			}
			continue
		}

		// Skip if we haven't encountered a section header yet
		if currentLED == nil {
			continue
		}

		// Parse key-value pairs
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Skip invalid lines
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "effect":
			currentLED.Effect, _ = strconv.Atoi(value)
		case "color1":
			// Remove "0x" prefix if it exists
			if len(value) > 2 && value[0:2] == "0x" {
				currentLED.Color1 = value[2:]
			} else {
				currentLED.Color1 = value
			}
		case "color2":
			// Remove "0x" prefix if it exists
			if len(value) > 2 && value[0:2] == "0x" {
				currentLED.Color2 = value[2:]
			} else {
				currentLED.Color2 = value
			}
		case "speed":
			currentLED.Speed, _ = strconv.Atoi(value)
		case "brightness":
			currentLED.Brightness, _ = strconv.Atoi(value)
		case "trigger":
			currentLED.Trigger, _ = strconv.Atoi(value)
		case "filename":
			currentLED.InternalName = value
		case "inbrightness":
			currentLED.InfoBrightness, _ = strconv.Atoi(value)
		}
	}

	// Don't forget to add the last LED
	if currentLED != nil {
		leds = append(leds, *currentLED)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	mappedLEDs := make(map[string]models.LED)

	for _, led := range leds {
		mappedLEDs[led.InternalName] = led
	}

	return mappedLEDs, nil
}

func WriteLEDSettings(leds []models.LED) error {
	file, err := os.Create(GetSettingsFile())
	if err != nil {
		return err
	}
	defer file.Close()

	for _, led := range leds {
		_, err := file.WriteString(led.FormatLedConfig())
		if err != nil {
			return err
		}

		_, err = file.WriteString("\n")
		if err != nil {
			return err
		}
	}

	return nil
}
