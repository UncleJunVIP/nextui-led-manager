package ui

import (
	gaba "github.com/UncleJunVIP/gabagool/pkg/gabagool"
	"nextui-led-control/functions"
	"nextui-led-control/models"
	"qlova.tech/sum"
)

type MainMenu struct {
}

func InitMainMenu() MainMenu {
	return MainMenu{}
}

func (m MainMenu) Name() sum.Int[models.ScreenName] {
	return models.ScreenNames.MainMenu
}

func (m MainMenu) Draw() (host interface{}, exitCode int, e error) {
	brickMenuItems := []gaba.MenuItem{
		{
			Text:               "Function Key 1",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           models.LED{DisplayName: "Function Key 1", InternalName: "f1"},
		},
		{
			Text:               "Function Key 2",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           models.LED{DisplayName: "Function Key 2", InternalName: "f2"},
		},
		{
			Text:               "Top Bar",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           models.LED{DisplayName: "Top Bar", InternalName: "m"},
		},
		{
			Text:               "L/R Triggers",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           models.LED{DisplayName: "L/R Triggers", InternalName: "lr"},
		},
	}

	smartProMenuItems := []gaba.MenuItem{
		{
			Text:               "Left Stick",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           models.LED{DisplayName: "Left Stick", InternalName: "l"},
		},
		{
			Text:               "Right Stick",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           models.LED{DisplayName: "Right Stick", InternalName: "r"},
		},
		{
			Text:               "TrimUI Center Logo",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           models.LED{DisplayName: "TrimUI Center Logo", InternalName: "m"},
		},
	}

	menuItems := brickMenuItems
	if !functions.IsBrick {
		menuItems = smartProMenuItems
	}

	options := gaba.DefaultListOptions("LED Manager", menuItems)
	options.EnableAction = true
	options.FooterHelpItems = []gaba.FooterHelpItem{
		{ButtonName: "B", HelpText: "Quit"},
		{ButtonName: "A", HelpText: "Select"},
	}

	selection, err := gaba.List(options)
	if err != nil {
		return nil, -1, err
	}

	if selection.IsSome() && selection.Unwrap().SelectedIndex != -1 {
		return selection.Unwrap().SelectedItem.Metadata, 0, nil
	}

	return nil, 2, nil
}
