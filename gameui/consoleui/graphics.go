package consoleui

import (
	"github.com/nsf/termbox-go"
	"github.com/xosmig/roguelike/core/geom"
	"fmt"
)

func (ui *consoleUi) drawMap() error {
	gameMap := ui.model.GetMap()

	for row := 0; row < gameMap.GetHeight(); row++ {
		for col := 0; col < gameMap.GetWidth(); col++ {
			var ch rune
			switch gameMap.Get(geom.Location{row, col}).Object.ModelName() {
			case "exit":
				ch = 'O'
			case "character":
				ch = '@'
			case "wall":
				ch = '#'
			case "empty":
				ch = ' '
			case "zombie":
				ch = 'z'
			case "item":
				ch = '$'
			default:
				return fmt.Errorf("invalid model name")
			}
			termbox.SetCell(col, row, rune(ch), 0, 0)
		}
	}
	return nil
}

func (ui *consoleUi) drawHealthBar() error {
	row := ui.mapBottom() + 1
	char := ui.model.GetCharacter()
	for i := 0; i < char.GetHP(); i++ {
		termbox.SetCell(i, row, '@', termbox.ColorRed|termbox.AttrBold, 0)
	}
	for i := char.GetHP(); i < char.GetMaxHP(); i++ {
		termbox.SetCell(i, row, '@', termbox.ColorCyan, 0)
	}

	return nil
}

func (ui *consoleUi) clear() {
	termbox.Clear(0, 0)
	ui.messagesLine = ui.mapBottom() + 5
}

func (ui *consoleUi) render() error {
	ui.clear()

	if err := ui.drawMap(); err != nil {
		return err
	}
	if err := ui.drawHealthBar(); err != nil {
		return err
	}

	return termbox.Flush()
}
