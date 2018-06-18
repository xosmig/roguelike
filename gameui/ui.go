package gameui

import "io"

// Ui interface represents a complete graphical application which can be run.
type Ui interface {
	io.Closer
	// Run actually runs the graphical application.
	// It blocks until the user decides to exit.
	// Safeness of invoking run multiple times depends on the particular implementation.
	Run() error
}
