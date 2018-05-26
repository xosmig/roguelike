package objects

import "github.com/xosmig/roguelike/core/geom"

type GameObject interface {
	Response(other GameObject)
	ModelName() string
}

type HasPosition interface {
	SetPosition(pos geom.Location)
	GetPosition() geom.Location
}

type MovableObject interface {
	GameObject
	HasPosition
	Interact(other GameObject)
}

type PositionData struct {
	pos geom.Location
}

func (p *PositionData) SetPosition(pos geom.Location) {
	p.pos = pos
}

func (p *PositionData) GetPosition() geom.Location {
	return p.pos
}

type wallObject struct{}

func (w wallObject) Response(other GameObject) {}
func (w wallObject) ModelName() string {
	return "wall"
}

var Wall GameObject = wallObject{}

type emptyObject struct{}

func (e emptyObject) Response(other GameObject) {}
func (e emptyObject) ModelName() string {
	return "empty"
}

var Empty GameObject = emptyObject{}
