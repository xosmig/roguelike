package geom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomDirection(tt *testing.T) {
	var res []Direction
	for i := 0; i < 100; i++ {
		res = append(res, RandomDirection())
	}

	allow := map[Direction]bool{
		Up:    true,
		Down:  true,
		Left:  true,
		Right: true,
	}
	met := make(map[Direction]bool)

	for _, dir := range res {
		assert.True(tt, allow[dir])
		met[dir] = true
	}

	assert.Equal(tt, len(allow), len(met))
}
