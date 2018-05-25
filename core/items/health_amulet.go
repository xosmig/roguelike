package items

import (
	"fmt"
	"github.com/xosmig/roguelike/core/character"
)

const maxHPBonus = 10

type HealthAmulet struct{}

func NewHealthAmulet() character.Item {
	return HealthAmulet{}
}

func (amulet HealthAmulet) Wear(character *character.Character) error {
	character.SetMaxHPPreserveNorm(character.GetMaxHP() + maxHPBonus)
	return nil
}

func (amulet HealthAmulet) TakeOff(character *character.Character) error {
	if character.GetMaxHP() <= maxHPBonus {
		return fmt.Errorf("this amulet keeps you alive")
	}
	character.SetMaxHPPreserveNorm(character.GetMaxHP() - maxHPBonus)
	return nil
}
