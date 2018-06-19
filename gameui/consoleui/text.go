package consoleui

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

// print prints the text in the game console.
// It's used to provide textual feedback to the user.
// Note that you will need to call termbox.Flush for user to see the printed text.
// Consider using messagef function.
func (ui *consoleUi) print(msg string) {
	for _, c := range msg {
		ui.draw(c, 0, 0)
	}
}

// println prints the message and ends the line.
// See documentation for print.
func (ui *consoleUi) println(msg string) {
	ui.print(msg)
	ui.nextLine()
}

// See documentation for print.
func (ui *consoleUi) printf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	ui.print(s)
}

// messagef prints the message to the user, ends the line and flushes the buffer
// so that the user can see the message immediately.
func (ui *consoleUi) messagef(format string, args ...interface{}) {
	ui.printf(format, args...)
	ui.nextLine()
	termbox.Flush()
}
