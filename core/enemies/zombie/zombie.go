package zombie

import (
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/unit"
	"github.com/xosmig/roguelike/core/state"
	"github.com/xosmig/roguelike/core/geom"
)

type Zombie struct {
	unit.UnitData
}

func New() *Zombie {
	return &Zombie{
		UnitData: unit.UnitData{
			MaxHP: 3,
			CurHP: 3,
			Team:  unit.TeamEvil,
		},
	}
}

func (zombie *Zombie) RecvDamage(dmg int, from unit.Unit) {
	unit.RecvDamageDefault(zombie, dmg, from)
}

func (zombie *Zombie) Interact(other objects.GameObject) {
	unit.InteractDefault(zombie, other)
}

func (zombie *Zombie) Response(other objects.GameObject) {
	// empty
}

func (zombie *Zombie) ModelName() string {
	return "zombie"
}

func (zombie *Zombie) DoAction(st state.GameState) {
	char := st.GetCharacter()
	if dir, ok := zombie.GetPosition().StepTo(char.GetPosition()); ok {
		st.TryMove(zombie, dir)
	} else {
		st.TryMove(zombie, geom.RandomDirection())
	}
}
