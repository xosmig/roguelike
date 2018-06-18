package consoleui

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/xosmig/roguelike/gamemodel"
	"github.com/xosmig/roguelike/gameui"
	"github.com/xosmig/roguelike/resources"
)

// consoleUi provides ascii graphics ui inside the terminal.
type consoleUi struct {
	mapName string
	model   gamemodel.GameModel

	curLine int
	linePos int
}

func (ui *consoleUi) draw(ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(ui.linePos, ui.curLine, ch, fg, bg)
	ui.linePos++
}

func (ui *consoleUi) nextLine() {
	ui.curLine++
	ui.linePos = 0
}

func (ui *consoleUi) emptyLine() {
	if ui.linePos != 0 {
		panic("Invalid use of consoleUi.emptyLine()")
	}
	ui.nextLine()
}

// New creates a new ascii graphics ui right inside the terminal.
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
