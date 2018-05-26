package state

import (
	"github.com/xosmig/roguelike/core/gamemap"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/unit"
	"github.com/xosmig/roguelike/core/character"
)

type GameState interface {
	GetMap() gamemap.GameMap
	GetCharacter() character.Character
	TryMove(obj objects.MovableObject, direction geom.Direction)
}

func RemoveDead(st GameState, u unit.Unit) bool {
	if !unit.IsAlive(u) {
		gamemap.Remove(st.GetMap(), u.GetPosition())
		return true
	}
	return false
}
