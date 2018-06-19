package consoleui

import (
	"github.com/nsf/termbox-go"
	"github.com/xosmig/roguelike/core/geom"
	"log"
)

func (ui *consoleUi) drawMap() {
	gameMap := ui.model.GetMap()

	for row := 0; row < gameMap.GetHeight(); row++ {
		for col := 0; col < gameMap.GetWidth(); col++ {
			var ch rune
			modelName := gameMap.Get(geom.Loc(row, col)).Object.ModelName()
			switch modelName {
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
				log.Printf("Error: unknown object model name: '%v'\n", modelName)
				ch = '?'
			}
			ui.draw(rune(ch), 0, 0)
		}
		ui.nextLine()
	}
}

func (ui *consoleUi) drawRules() {
	ui.println("Use arrows to move, Ctrl+C to exit")
}

func (ui *consoleUi) drawHealthBar() {
	char := ui.model.GetCharacter()

	for i := 0; i < char.GetHP(); i++ {
		ui.draw('@', termbox.ColorRed|termbox.AttrBold, 0)
	}
	for i := char.GetHP(); i < char.GetMaxHP(); i++ {
		ui.draw('@', termbox.ColorCyan, 0)
	}

	ui.nextLine()
}

func (ui *consoleUi) drawInventory() {
	char := ui.model.GetCharacter()

	if len(char.Inventory()) == 0 {
		ui.emptyLine()
		ui.emptyLine()
		ui.emptyLine()
		return
	}

	ui.println("Use Ctrl + A/S/D to wear or take off items")
	for i := range char.Inventory() {
		ui.draw('A'+rune(i), 0, 0)
	}
	ui.nextLine()

	for _, item := range char.Inventory() {
		var ch rune
		switch item.IconName() {
		case "health_amulet":
			ch = 'H'
		default:
			log.Printf("Error: unknown item icon name: '%v'\n", item.IconName())
			ch = '?'
		}

		var attr = termbox.ColorCyan
		if char.Wearing() == item {
			attr = termbox.ColorRed | termbox.AttrBold
		}

		ui.draw(ch, attr, 0)
	}
	ui.nextLine()
}

func (ui *consoleUi) clear() {
	termbox.Clear(0, 0)
	ui.curLine = 0
}

// render clears the screen and renders the next frame.
func (ui *consoleUi) render() error {
	ui.clear()

	ui.drawMap()
	ui.emptyLine()
	ui.drawRules()
	ui.drawHealthBar()
	ui.emptyLine()
	ui.drawInventory()

	ui.emptyLine()
	ui.emptyLine()
	// messages will be displayed bellow

	return termbox.Flush()
}
