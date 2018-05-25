package consoleui

import (
	"github.com/nsf/termbox-go"
	"github.com/xosmig/roguelike/core/geom"
	"log"
	"fmt"
	"github.com/xosmig/roguelike/gamemodel/status"
)

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
	ui.messagef("Press 'Ctrl+C' to exit, or 'Ctrl+R' to restart")

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

func (ui *consoleUi) Run() error {
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
			ui.messagef("You lost :(")
			return ui.restartOrExit()
		case status.Victory:
			ui.messagef("You won [^_^]")
			return ui.restartOrExit()
		}

		key, finish := ui.getKeyForAction(actions)
		if finish {
			break gameLoop
		}

		actions[key]()
	}

	return nil
}
