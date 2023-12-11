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
	tests := []struct {
		name      string
		increment int
		expected  int
	}{
		{
			name:      "1",
			increment: 1,
			expected:  374,
		}, {
			name:      "10",
			increment: 9,
			expected:  1030,
		}, {
			name:      "100",
			increment: 99,
			expected:  8410,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.name), func(t *testing.T) {
			res := ExpandGalaxy(input, tt.increment)
			assert.Equal(t, tt.expected, res)
		})
	}
}
