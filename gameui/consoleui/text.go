package consoleui

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

func printTb(x, y int, fg, bg termbox.Attribute, msg string) {

}

func (ui *consoleUi) print(msg string) {
	for _, c := range msg {
		ui.draw(c, 0, 0)
	}
}

func (ui *consoleUi) println(msg string) {
	ui.print(msg)
	ui.nextLine()
}

func (ui *consoleUi) printf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	ui.print(s)
}


func (ui *consoleUi) messagef(format string, args ...interface{}) {
	ui.printf(format, args...)
	ui.nextLine()
	termbox.Flush()
}
