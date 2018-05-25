package character

type Item interface {
	Wear(character *Character) error
	TakeOff(character *Character) error
}
