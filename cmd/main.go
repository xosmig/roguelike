package main

import (
	"github.com/xosmig/roguelike/gameui/consoleui"
)

func main() {
	//logFile, err := os.Create("log.txt")
	//if err != nil {
	//	panic("Error creating log file")
	//}
	//log.SetOutput(logFile)

	ui, err := consoleui.New("builtin/maps/example")
	if err != nil {
		panic(err)
	}
	err = ui.Run()
	if err != nil {
		panic(err)
	}
}
