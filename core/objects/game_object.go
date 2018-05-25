package objects

type GameObject interface {
	Interact(other GameObject)
	Response(other GameObject)
	ModelName() string
}

type wallObject struct{}

func (w wallObject) Interact(other GameObject) {}
func (w wallObject) Response(other GameObject) {}
func (w wallObject) ModelName() string {
	return "wall"
}

var Wall GameObject = wallObject{}

type emptyObject struct{}

func (e emptyObject) Interact(other GameObject) {}
func (e emptyObject) Response(other GameObject) {}
func (e emptyObject) ModelName() string {
	return "empty"
}

var Empty GameObject = emptyObject{}

type ExitHandler interface {
	OnExit()
}

type Exit struct {
	handler ExitHandler
}

func NewExit(handler ExitHandler) Exit {
	return Exit{handler}
}

func (e Exit) Interact(other GameObject) {}

func (e Exit) Response(other GameObject) {
	e.handler.OnExit()
}

func (e Exit) ModelName() string {
	return "exit"
}
