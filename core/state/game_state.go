package state

import (
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/core/gamemap"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/unit"
)

// GameState used by game objects to access game state from callbacks such as `DoAction`
type GameState interface {
	GetMap() gamemap.GameMap
	GetCharacter() character.Character
	// Tries to move the object to the given direction.
	// If there is another object, `ojb.Interact` and `otherObject.Response` are called
	// and no movement happens unless the other object destroys itself in it's response implementation.
	TryMove(obj objects.MovableObject, direction geom.Direction)
}

// RemoveDead removes unit from map if it's not alive
func RemoveDead(st GameState, u unit.Unit) bool {
	if !unit.IsAlive(u) {
		gamemap.Remove(st.GetMap(), u.GetPosition())
		return true
	}
	return false
}
