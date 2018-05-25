package state

import (
	"github.com/xosmig/roguelike/core/gamemap"
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/core/unit"
	"github.com/xosmig/roguelike/core/geom"
)

type GameState interface {
	GetMap() gamemap.GameMap
	GetCharacter() *character.Character
	TryMove(unit unit.Unit, direction geom.Direction)
}
