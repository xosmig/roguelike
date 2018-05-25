package consoleui

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

func printTb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printfTb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	printTb(x, y, fg, bg, s)
}

func (ui *consoleUi) messagef(format string, args ...interface{}) {
	printfTb(0, ui.messagesLine, 0, 0, format, args...)
	ui.messagesLine++
	termbox.Flush()
}

func (ui *consoleUi) clearMessages() {
	ui.messagesLine = ui.mapBottom() + 5
}
