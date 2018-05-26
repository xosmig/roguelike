package geom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocation_StepTo(t *testing.T) {
	t.Run("True_Right", func(tt *testing.T) {
		dir, neighbour := Loc(5, 5).StepTo(Loc(5, 6))
		assert.True(tt, neighbour)
		assert.Equal(tt, Right, dir)
	})

	t.Run("True_Down", func(tt *testing.T) {
		dir, neighbour := Loc(5, 5).StepTo(Loc(6, 5))
		assert.True(tt, neighbour)
		assert.Equal(tt, Down, dir)
	})

	t.Run("False_Diagonal", func(tt *testing.T) {
		dir, neighbour := Loc(5, 5).StepTo(Loc(4, 4))
		assert.False(tt, neighbour)
		assert.Equal(tt, Nowhere, dir)
	})

	t.Run("False_Horizontal", func(tt *testing.T) {
		dir, neighbour := Loc(5, 5).StepTo(Loc(5, 7))
		assert.False(tt, neighbour)
		assert.Equal(tt, Nowhere, dir)
	})

	t.Run("True_SameLocation", func(tt *testing.T) {
		dir, neighbour := Loc(5, 5).StepTo(Loc(5, 5))
		assert.True(tt, neighbour)
		assert.Equal(tt, Nowhere, dir)
	})
}
