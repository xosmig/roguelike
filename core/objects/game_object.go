package objects

type GameObject interface {
	Interact(other GameObject)
	Response(other GameObject)
	ModelName() string
}

type HasPosition interface {
	SetPosition(pos Location)
	GetPosition() Location
}

type PositionData struct {
	Position Location
}

func (p *PositionData) SetPosition(pos Location) {
	p.Position = pos
}

func (p *PositionData) GetPosition() Location {
	return p.Position
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
