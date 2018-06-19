// Tests have to be placed in a separate package to avoid dependency cycle between `unit` and `mock_unit`
package tests

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xosmig/roguelike/core/objects/mock_objects"
	"github.com/xosmig/roguelike/core/unit"
	"github.com/xosmig/roguelike/core/unit/mock_unit"
	"testing"
)

func TestInteractDefault(t *testing.T) {
	t.Run("AgainstNotUnit", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		u := mock_unit.NewMockUnit(ctrl)
		target := mock_objects.NewMockGameObject(ctrl)

		unit.InteractDefault(u, target)
	})

	t.Run("AgainstEnemy", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		u := mock_unit.NewMockUnit(ctrl)
		u.EXPECT().GetTeam().MinTimes(1).Return(unit.TeamGood)
		target := mock_unit.NewMockUnit(ctrl)
		target.EXPECT().GetTeam().MinTimes(1).Return(unit.TeamEvil)
		target.EXPECT().RecvDamage(1, u)

		unit.InteractDefault(u, target)
	})

	t.Run("AgainstEnemy", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		u := mock_unit.NewMockUnit(ctrl)
		u.EXPECT().GetTeam().MinTimes(1).Return(unit.TeamGood)
		target := mock_unit.NewMockUnit(ctrl)
		target.EXPECT().GetTeam().MinTimes(1).Return(unit.TeamGood)

		unit.InteractDefault(u, target)
	})
}

func TestIsAlive(t *testing.T) {
	t.Run("true", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		u := mock_unit.NewMockUnit(ctrl)
		u.EXPECT().GetHP().Return(3)
		assert.True(tt, unit.IsAlive(u))
	})

	t.Run("false", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		u := mock_unit.NewMockUnit(ctrl)
		u.EXPECT().GetHP().Return(0)
		assert.False(tt, unit.IsAlive(u))
	})
}
