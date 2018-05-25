package gameui

import "io"

type Ui interface {
	io.Closer
	Run() error
}
