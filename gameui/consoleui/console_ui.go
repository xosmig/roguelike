package consoleui

import (
	"fmt"
	"github.com/xosmig/roguelike/gamemodel"
	"github.com/nsf/termbox-go"
	"github.com/xosmig/roguelike/gameui"
	"github.com/xosmig/roguelike/resources"
	"github.com/xosmig/roguelike/core/geom"
	"log"
	"github.com/xosmig/roguelike/gamemodel/status"
)

type consoleUi struct {
	mapName string
	model   gamemodel.GameModel
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

func (ui *consoleUi) render() error {
	termbox.Clear(0, 0)
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
			default:
				return fmt.Errorf("invalid model name")
			}
			termbox.SetCell(col, row, rune(ch), 0, 0)
		}
	}

	return termbox.Flush()
}

func (ui *consoleUi) getKeyForAction(actions map[termbox.Key]func()) (key termbox.Key, finish bool) {
	var ev termbox.Event
	for {
		ev = termbox.PollEvent()
		log.Printf("FOO: %v %v", ev.Type, ev.Key)
		if ev.Type != termbox.EventKey {
			continue
		}
		if ev.Key == termbox.KeyCtrlC {
			log.Println("Ctrl+C is pressed")
			return ev.Key, true
		}
		if _, present := actions[ev.Key]; present {
			return ev.Key, false
		} else {
			log.Printf("Debug: Invalid command key: %v\n", ev.Key)
			continue
		}
	}
}

func (ui *consoleUi) restartOrExit() error {
	fmt.Println("Press 'Ctrl+C' to exit, or 'Ctrl+R' to restart")

	actions := map[termbox.Key]func(){
		termbox.KeyCtrlR: nil,
	}

	_, finish := ui.getKeyForAction(actions)
	if finish {
		log.Println("Exit requested")
		return nil
	}

	ui.reloadGameModel()
	return ui.Run()
}

func (ui *consoleUi) resetText() {
	termbox.SetCursor(0, ui.model.GetMap().GetHeight()+5)
	termbox.Sync()
}

func (ui *consoleUi) printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (ui *consoleUi) Run() error {
	ui.resetText()

	actions := map[termbox.Key]func(){
		termbox.KeyArrowUp:    func() { ui.model.DoMove(geom.Up) },
		termbox.KeyArrowDown:  func() { ui.model.DoMove(geom.Down) },
		termbox.KeyArrowLeft:  func() { ui.model.DoMove(geom.Left) },
		termbox.KeyArrowRight: func() { ui.model.DoMove(geom.Right) },
	}

gameLoop:
	for {
		err := ui.render()
		if err != nil {
			return fmt.Errorf("while rendering: %v", err)
		}

		switch ui.model.Status() {
		case status.Continue:
			// continue
		case status.Defeat:
			ui.printf("You lost :(\n")
			return ui.restartOrExit()
		case status.Victory:
			ui.printf("You won [^_^]\n")
			return ui.restartOrExit()
		}

		key, finish := ui.getKeyForAction(actions)
		if finish {
			break gameLoop
		}

		// clear all text messages
		ui.resetText()

		actions[key]()
	}

	return nil
}
