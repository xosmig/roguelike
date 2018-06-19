package gamemodel

import (
	"github.com/stretchr/testify/assert"
	"github.com/xosmig/roguelike/core/enemies/zombie"
	"github.com/xosmig/roguelike/core/geom"
	"github.com/xosmig/roguelike/core/items"
	"github.com/xosmig/roguelike/core/objects"
	"github.com/xosmig/roguelike/core/unit"
	"github.com/xosmig/roguelike/gamemodel/status"
	"github.com/xosmig/roguelike/resources"
	"testing"
)

var _ objects.HasPosition = &exit{}

var simpleMap = "4 5\n" +
	"#####\n" +
	"#@$.#\n" +
	"#z.O#\n" +
	"#####\n" +
	""

var resourceLoader = resources.NewInMemoryLoader()

func init() {
	resourceLoader.AddString("simple_map", simpleMap)
}

func newModel(tt *testing.T) GameModel {
	model, err := New(resourceLoader, "simple_map")
	assert.NoError(tt, err)
	return model
}

func assertPosition(tt *testing.T, model GameModel, pos geom.Location) {
	assert.Equal(tt, model.GetCharacter().GetPosition(), pos)
}

func findZombie(model GameModel) *zombie.Zombie {
	return model.GetMap().Get(geom.Loc(2, 1)).Object.(*zombie.Zombie)
}

func findAmulet(model GameModel) *items.ItemObject {
	return model.GetMap().Get(geom.Loc(1, 2)).Object.(*items.ItemObject)
}

func TestNew(tt *testing.T) {
	model := newModel(tt)
	assert.NotNil(tt, model.GetMap())
	assert.Equal(tt, status.Continue, model.Status())

	char := model.GetCharacter()
	assert.NotNil(tt, char)
	assert.Equal(tt, geom.Loc(1, 1), char.GetPosition())
	assert.Equal(tt, char.GetMaxHP(), char.GetHP())

	z := findZombie(model)
	assert.NotNil(tt, z)
	assert.Equal(tt, geom.Loc(2, 1), z.GetPosition())
	assert.Equal(tt, z.GetMaxHP(), z.GetHP())

	a := findAmulet(model)
	assert.NotNil(tt, a)
	assert.Equal(tt, geom.Loc(1, 2), a.GetPosition())
}

func TestGameModel_DoMove(t *testing.T) {
	t.Run("ToEmptyCell", func(tt *testing.T) {
		model := newModel(tt)
		model.DoMove(geom.Right)
		assertPosition(tt, model, geom.Loc(1, 2))
	})

	t.Run("ToWall", func(tt *testing.T) {
		model := newModel(tt)
		model.DoMove(geom.Up)
		assertPosition(tt, model, geom.Loc(1, 1))
	})

	t.Run("Attack", func(tt *testing.T) {
		model := newModel(tt)
		z := findZombie(model)
		model.DoMove(geom.Down)
		assert.Equal(tt, z.GetMaxHP()-1, z.GetHP())
	})

	t.Run("Stay", func(tt *testing.T) {
		model := newModel(tt)
		model.DoMove(geom.Nowhere)
		assertPosition(tt, model, geom.Loc(1, 1))
	})

	t.Run("ZombieAttacks", func(tt *testing.T) {
		model := newModel(tt)
		model.DoMove(geom.Nowhere)
		char := model.GetCharacter()
		assert.Equal(tt, char.GetMaxHP()-1, char.GetHP())
	})
}

func TestGameModel_Scenario(t *testing.T) {
	t.Run("Victory", func(tt *testing.T) {
		model := newModel(tt)

		model.DoMove(geom.Right)
		model.DoMove(geom.Right)

		assert.Equal(tt, status.Continue, model.Status())

		model.DoMove(geom.Down)

		assert.Equal(tt, status.Victory, model.Status())
		assertPosition(tt, model, geom.Loc(2, 3))
	})

	t.Run("DeathAndDefeat", func(tt *testing.T) {
		model := newModel(tt)
		pos := model.GetCharacter().GetPosition()

		model.DoMove(geom.Nowhere)
		model.DoMove(geom.Nowhere)

		assert.Equal(tt, status.Continue, model.Status())
		assert.True(tt, unit.IsAlive(model.GetCharacter()))
		assert.NotEqual(tt, objects.Empty, model.GetMap().Get(pos).Object)

		model.DoMove(geom.Nowhere)

		assert.Equal(tt, status.Defeat, model.Status())
		assert.False(tt, unit.IsAlive(model.GetCharacter()))
		assert.Equal(tt, objects.Empty, model.GetMap().Get(pos).Object)
	})
}
