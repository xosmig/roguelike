package character

import (
	"fmt"
)

const maxHPBonus = 10

type HealthAmulet struct{}

func (amulet *HealthAmulet) Wear(character *Character) error {
	character.SetMaxHPPreserveNorm(character.GetMaxHP() + maxHPBonus)
	return nil
}

func (amulet *HealthAmulet) TakeOff(character *Character) error {
	if character.GetMaxHP() <= maxHPBonus {
		return fmt.Errorf("this amulet keeps you alive")
	}
	character.SetMaxHPPreserveNorm(character.GetMaxHP() - maxHPBonus)
	return nil
}
