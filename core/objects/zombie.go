package objects

type Zombie struct {
	UnitData
}

func (zombie *Zombie) RecvDamage(dmg int, from Unit) {
	RecvDamageDefault(zombie, dmg, from)
}

func (zombie *Zombie) Die(from Unit) {
	// TODO
}

func (zombie *Zombie) Interact(other GameObject) {
	UnitInteractDefault(zombie, other)
}

func (zombie *Zombie) Response(other GameObject) {
	zombie.Interact(other)
}

func (zombie *Zombie) ModelName() string {
	return "zombie"
}
