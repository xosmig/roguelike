package ai

import "github.com/xosmig/roguelike/core/state"

// Describes the callback called by game model every turn
type Actionable interface {
	DoAction(st state.GameState)
}
