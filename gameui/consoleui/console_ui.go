package consoleui

import (
	"fmt"
	"github.com/xosmig/roguelike/gamemodel"
	"github.com/nsf/termbox-go"
	"github.com/xosmig/roguelike/gameui"
	"github.com/xosmig/roguelike/resources"
)

type consoleUi struct {
	mapName      string
	model        gamemodel.GameModel
	messagesLine int
}

func New(mapName string) (gameui.Ui, error) {
	ui := &consoleUi{mapName: mapName}

	if err := termbox.Init(); err != nil {
		return nil, fmt.Errorf("while initializing termbox: %v", err)
	}

	if err := ui.reloadGameModel(); err != nil {
		return nil, err
	}
	return ui, nil
}

func (ui *consoleUi) reloadGameModel() error {
	model, err := gamemodel.New(resources.BuiltinLoader, ui.mapName)
	if err != nil {
		return fmt.Errorf("while initializing game model: %v", err)
	}
	ui.model = model
	return nil
}

func (ui *consoleUi) Close() error {
	termbox.Close()
	return nil
}

func (ui *consoleUi) mapBottom() int {
	return ui.model.GetMap().GetHeight()
}
