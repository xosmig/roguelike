package main

import (
	"github.com/xosmig/roguelike/gameui/consoleui"
	"os"
	"log"
	"flag"
	"fmt"
)

func main() {
	logFile := flag.String("logfile", "", "File to write log to")
	flag.Parse()

	if *logFile == "" {
		*logFile = os.DevNull
	}
	f, err := os.Create(*logFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating log file")
		os.Exit(2)
	}
	log.SetOutput(f)

	ui, err := consoleui.New("builtin/maps/example")
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	err = ui.Run()
	if err != nil {
		panic(err)
	}
}
