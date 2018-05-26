package state

import (
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/core/gamemap"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/unit"
)

// Used by game objects to access game state from callbacks such as `DoAction`
type GameState interface {
	GetMap() gamemap.GameMap
	GetCharacter() character.Character
	TryMove(obj objects.MovableObject, direction geom.Direction)
}

// Remove unit from map if it's not alive
func RemoveDead(st GameState, u unit.Unit) bool {
	if !unit.IsAlive(u) {
		gamemap.Remove(st.GetMap(), u.GetPosition())
		return true
	}
	return false
}
