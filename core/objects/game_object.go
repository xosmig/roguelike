package objects

import "github.com/xosmig/roguelike/core/geom"

// GameObject represents an object on map.
// Multiple map cells may share the same GameObject.
// For example, objects.Wall is used for every wall cell on map.
type GameObject interface {
	// Response is called when another object is trying to interact with this object
	// by moving to its position on map.
	Response(other GameObject)

	// ModelName returns the name of the graphical representation of the object.
	// It should always return the same value for the same object.
	// The returned value must be one of the values, defined by the user interface used.
	ModelName() string
}

// HasPosition should be implemented by the objects which want to track their positions.
// Note that it doesn't make sense to place the same instance of GameObject which implements
// HasPosition interface to more than one map cell.
//
// It's advised to implement this interface by embedding PositionData structure.
type HasPosition interface {
	// SetPosition is called by the game model each time when the object is moved to a new position.
	SetPosition(pos geom.Location)
	GetPosition() geom.Location
}

// MovableObject represents a GameObject which might try to move to a new position,
// and hence cause an interaction.
type MovableObject interface {
	GameObject
	HasPosition
	// Interact is called when this GameObject is trying to move to a cell, where
	// another object is placed.
	// Call to Interact method of this object is followed by a call to
	// Response method of the other object.
	Interact(other GameObject)
}

// PositionData provides the basic implementation of HasPosition interface.
type PositionData struct {
	pos geom.Location
}

// SetPosition updates the knowledge of the object about its current position.
// See documentation for HasPosition interface.
func (p *PositionData) SetPosition(pos geom.Location) {
	p.pos = pos
}

// GetPosition return the argument of the last call to SetPosition.
func (p *PositionData) GetPosition() geom.Location {
	return p.pos
}

type wallObject struct{}

func (w wallObject) Response(other GameObject) {}
func (w wallObject) ModelName() string {
	return "wall"
}

// Wall object is used to represent a wall on a game map.
var Wall GameObject = wallObject{}

type emptyObject struct{}

func (e emptyObject) Response(other GameObject) {}
func (e emptyObject) ModelName() string {
	return "empty"
}

// Empty is a special instance of GameObject It represents an empty cell.
var Empty GameObject = emptyObject{}
