package ui

import (
	gaba "github.com/UncleJunVIP/gabagool/pkg/gabagool"
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
	menuItems := []gaba.MenuItem{
		{
			Text:               "Function Key 1",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           "F1",
		},
		{
			Text:               "Function Key 2",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           "F2",
		},
		{
			Text:               "Top Bar",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           "Topbar",
		},
		{
			Text:               "L/R Triggers",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           "Triggers",
		},
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

	if selection.IsSome() {
		return selection.Unwrap().SelectedItem.Metadata, 0, nil
	}

	return nil, 2, nil
}
