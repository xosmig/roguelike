package items

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xosmig/roguelike/core/character/mock_character"
	"testing"
)

func TestHealthAmulet(t *testing.T) {
	amulet := NewHealthAmulet()

	t.Run("IncreasesHp", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		char := mock_character.NewMockCharacter(ctrl)

		char.EXPECT().GetHP().MinTimes(1).Return(3)
		char.EXPECT().SetHP(6)
		char.EXPECT().GetMaxHP().MinTimes(1).Return(10)
		char.EXPECT().SetMaxHP(13)
		assert.NoError(tt, amulet.Wear(char))
	})

	t.Run("DecreasesHp", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		char := mock_character.NewMockCharacter(ctrl)

		char.EXPECT().GetHP().MinTimes(1).Return(6)
		char.EXPECT().SetHP(3)
		char.EXPECT().GetMaxHP().MinTimes(1).Return(13)
		char.EXPECT().SetMaxHP(10)
		assert.NoError(tt, amulet.TakeOff(char))
	})

	t.Run("DoesntKill", func(tt *testing.T) {
		ctrl := gomock.NewController(tt)
		char := mock_character.NewMockCharacter(ctrl)

		char.EXPECT().GetHP().MinTimes(1).Return(3)
		assert.Error(tt, amulet.TakeOff(char))
	})
}
