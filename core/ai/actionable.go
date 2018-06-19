package ai

import "github.com/xosmig/roguelike/core/state"

// Actionable describes the callback called by game model every turn
type Actionable interface {
	DoAction(st state.GameState)
}
