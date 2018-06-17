package items

import (
	"github.com/xosmig/roguelike/core/character"
	"github.com/xosmig/roguelike/core/gamemap"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/state"
)

// ItemObject represents a collectible object on map.
// It behaves like a wall for all game objects except the character.
type ItemObject struct {
	objects.PositionData
	item character.Item
	st   state.GameState
}

// NewItemObject creates a new ItemObject which contains the provided item.
func NewItemObject(st state.GameState, item character.Item) *ItemObject {
	return &ItemObject{
		st:   st,
		item: item,
	}
}

// Response ignores any attempt to interact from any game object but the character.
// When the interaction is caused by the character, the ItemObject disappears
// (so, the character proceeds its movement),
// and the item appears in the character's inventory.
func (obj *ItemObject) Response(other objects.GameObject) {
	if char, ok := other.(character.Character); ok {
		char.AddItem(obj.item)
		gamemap.Remove(obj.st.GetMap(), obj.GetPosition())
	}
}

// ModelName returns the name of the graphical representation of the object.
// See the documentation for object.GameObject.ModelName.
func (obj *ItemObject) ModelName() string {
	return "item"
}
