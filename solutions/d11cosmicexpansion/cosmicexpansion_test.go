package d11cosmicexpansion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandGalaxy(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	a := ExpandGalaxy(input)
	assert.Equal(t, 374, a)
}
