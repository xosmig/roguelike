
package gamemodel
/*
import "github.com/xosmig/roguelike/gamemodel/objects"

type ObjectsSet struct {
	Objects []objects.GameObject
}

func (s *ObjectsSet) RemoveDeadObjects() {
	j := 0
	for i := 0; i < len(s.Objects); i++ {
		if s.Objects[i].IsAlive() {
			s.Objects[j] = s.Objects[i]
			j++
		}
	}
	s.Objects = s.Objects[:j]
}

func (s *ObjectsSet) Del(object objects.GameObject) {
	for i, item := range s.Objects {
		if item == object {
			s.Objects[i] = s.Objects[len(s.Objects) - 1]
			s.Objects = s.Objects[:len(s.Objects) - 1]
			return
		}
	}
}

func (s*ObjectsSet) Add(object objects.GameObject) {
	s.Objects = append(s.Objects, object)
}
*/
