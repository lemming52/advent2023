package d16thefloorwillbelava

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnergise(t *testing.T) {
	input := []string{
		".|...\\....",
		"|.-.\\.....",
		".....|-...",
		"........|.",
		"..........",
		".........\\",
		"..../.\\\\..",
		".-.-/..|..",
		".|....-|.\\",
		"..//.|....",
	}
	a := Energise(input)
	assert.Equal(t, 46, a)
}
