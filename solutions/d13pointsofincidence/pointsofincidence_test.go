package d13pointsofincidence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPlates(t *testing.T) {
	input := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
		"",
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}
	a := CheckPlates(input)
	assert.Equal(t, 405, a)
}
