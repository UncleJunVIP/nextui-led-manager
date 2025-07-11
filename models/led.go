package models

import (
	"fmt"
	"strings"
)

type LED struct {
	DisplayName    string
	InternalName   string
	Color1         string
	Color2         string
	Brightness     int
	InfoBrightness int
	Effect         int
	Speed          int
	Trigger        int
}

func (l *LED) FormatLedConfig() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("[%s]\n", l.DisplayName))
	sb.WriteString(fmt.Sprintf("effect=%d\n", l.Effect))
	sb.WriteString(fmt.Sprintf("color1=0x%s\n", l.Color1))
	sb.WriteString(fmt.Sprintf("color2=0x%s\n", l.Color2))
	sb.WriteString(fmt.Sprintf("speed=%d\n", l.Speed))
	sb.WriteString(fmt.Sprintf("brightness=%d\n", l.Brightness))
	sb.WriteString(fmt.Sprintf("trigger=%d\n", l.Trigger))
	sb.WriteString(fmt.Sprintf("filename=%s\n", l.InternalName))
	sb.WriteString(fmt.Sprintf("inbrightness=%d\n", l.InfoBrightness))

	return sb.String()
}
