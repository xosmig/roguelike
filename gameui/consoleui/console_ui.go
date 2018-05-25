package consoleui

import (
	"fmt"
	"github.com/xosmig/roguelike/gamemodel"
	"github.com/nsf/termbox-go"
	"github.com/xosmig/roguelike/gameui"
	"github.com/xosmig/roguelike/resources"
	"github.com/xosmig/roguelike/core/objects"
)

type consoleUi struct {
	model gamemodel.GameModel
}

func New(mapName string) (gameui.Ui, error) {
	ui := new(consoleUi)
	var err error
	ui.model, err = gamemodel.New(resources.BuiltinLoader, mapName, ui)
	if err != nil {
		return nil, fmt.Errorf("while initializing game model: %v", err)
	}
	return ui, nil
}

func (ui *consoleUi) OnExit() {
	// TODO
}

func (ui *consoleUi) render() error {
	termbox.Clear(0, 0)
	gameMap := ui.model.GetMap()

	for row := 0; row < gameMap.GetHeight(); row++ {
		for col := 0; col < gameMap.GetWidth(); col++ {
			var ch byte
			switch gameMap.Get(objects.Location{row, col}).Object.ModelName() {
			case "exit":
				ch = 'O'
			case "character":
				ch = '@'
			case "wall":
				ch = '#'
			case "empty":
				ch = ' '
			default:
				return fmt.Errorf("invalid model name")
			}
			termbox.SetCell(col, row, rune(ch), 0, 0)
		}
	}

	return termbox.Flush()
}

func (ui *consoleUi) Run() error {
	err := termbox.Init()
	if err != nil {
		return fmt.Errorf("while initializing termbox: %v", err)
	}
	defer termbox.Close()

	termbox.SetCursor(0, ui.model.GetMap().GetHeight()+5)

loop:
	for {
		err := ui.render()
		if err != nil {
			return fmt.Errorf("while rendering: %v", err)
		}

		ev := termbox.Event{Type: termbox.EventNone}
		for ev.Type != termbox.EventKey {
			ev = termbox.PollEvent()
		}
		// clear all text messages
		termbox.Sync()

		switch ev.Key {
		case termbox.KeyArrowUp:
			ui.model.DoMove(objects.Up)
		case termbox.KeyArrowDown:
			ui.model.DoMove(objects.Down)
		case termbox.KeyArrowLeft:
			ui.model.DoMove(objects.Left)
		case termbox.KeyArrowRight:
			ui.model.DoMove(objects.Right)
		case termbox.KeyCtrlC:
			break loop
		}
	}

	return nil
}
