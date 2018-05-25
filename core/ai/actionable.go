package ai

import "github.com/xosmig/roguelike/core/state"

type Actionable interface {
	DoAction(st state.GameState)
}
