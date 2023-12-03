package d03gearratios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartNumbers(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	a, b := PartNumbers(input)
	assert.Equal(t, 4361, a)
	assert.Equal(t, 467835, b)
}
