package d10pipemaze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraversePipe(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
		x, y     int
	}{
		{
			name: "simple",
			input: []string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			},
			expected: 4,
			x:        1,
			y:        1,
		}, {
			name: "longer",
			input: []string{
				"..F7.",
				".FJ|.",
				"SJ.L7",
				"|F--J",
				"LJ...",
			},
			expected: 8,
			x:        0,
			y:        2,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.name), func(t *testing.T) {
			a := TraversePipe(tt.input, tt.x, tt.y)
			assert.Equal(t, tt.expected, a)
		})
	}
}
