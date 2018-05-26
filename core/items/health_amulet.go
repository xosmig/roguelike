package items

import (
	"fmt"
	"github.com/xosmig/roguelike/core/character"
)

const amuletHPBonus = 3

type HealthAmulet struct{}

func NewHealthAmulet() character.Item {
	return HealthAmulet{}
}

func (amulet HealthAmulet) Wear(character character.Character) error {
	character.SetMaxHP(character.GetMaxHP() + amuletHPBonus)
	character.SetHP(character.GetHP() + amuletHPBonus)
	return nil
}

func (amulet HealthAmulet) TakeOff(character character.Character) error {
	if character.GetHP() <= amuletHPBonus {
		return fmt.Errorf("this amulet keeps you alive")
	}

	character.SetHP(character.GetHP() - amuletHPBonus)
	character.SetMaxHP(character.GetMaxHP() - amuletHPBonus)
	return nil
}

func (amulet HealthAmulet) IconName() string {
	return "health_amulet"
}
