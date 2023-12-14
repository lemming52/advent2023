package d14parabolicreflectordish

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFocusDish(t *testing.T) {
	input := []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
	a := FocusDish(input)
	assert.Equal(t, 136, a)
}
