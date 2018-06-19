package consoleui

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/gamemodel/status"
	"log"
)

// getKeyForAction accepts the set of interesting keys and waits for one of these keys to be pressed,
// or for a cancellation request from the user (Ctrl+C pressed).
// It returns the pressed key and a boolean flag, indicating whether the exit was requested or not.
func (ui *consoleUi) getKeyForAction(actions map[termbox.Key]func()) (key termbox.Key, finish bool) {
	var ev termbox.Event
	for {
		ev = termbox.PollEvent()
		if ev.Type != termbox.EventKey {
			continue
		}
		if ev.Key == termbox.KeyCtrlC {
			log.Println("Ctrl+C is pressed")
			return ev.Key, true
		}
		if _, present := actions[ev.Key]; present {
			return ev.Key, false
		}

		log.Printf("Debug: Invalid command key: %v\n", ev.Key)
	}
}

// restartOrExit blocks until the user presses Ctrl+C (in this case it just returns nil),
// or Ctrl+R (in this case it restarts the game via recursive call to Run method).
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

// Run does the read-execute-print-loop.
func (ui *consoleUi) Run() error {
	var afterRender []func()
	// delay delays the given function execution so that it is executed after rendering.
	// It's useful to adjust the rendered picture.
	// For example, by printing a message.
	delay := func(f func()) {
		afterRender = append(afterRender, f)
	}

	accessItem := func(idx int) {
		err := ui.model.GetCharacter().WearOrTakeOff(idx)
		if err != nil {
			delay(func() { ui.messagef("inventory error: %v", err) })
		}
	}

	actions := map[termbox.Key]func(){
		termbox.KeyArrowUp:    func() { ui.model.DoMove(geom.Up) },
		termbox.KeyArrowDown:  func() { ui.model.DoMove(geom.Down) },
		termbox.KeyArrowLeft:  func() { ui.model.DoMove(geom.Left) },
		termbox.KeyArrowRight: func() { ui.model.DoMove(geom.Right) },
		termbox.KeyCtrlA:      func() { accessItem(0) },
		termbox.KeyCtrlS:      func() { accessItem(1) },
		termbox.KeyCtrlD:      func() { accessItem(2) },
	}

gameLoop:
	for {
		err := ui.render()
		if err != nil {
			return fmt.Errorf("while rendering: %v", err)
		}

		for _, f := range afterRender {
			f()
		}
		afterRender = nil

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
