package d08hauntedwasteland

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrossDesert(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		input    []string
	}{
		{
			name:     "short",
			expected: 2,
			input: []string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},
		}, {
			name:     "long",
			expected: 6,
			input: []string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.name), func(t *testing.T) {
			a := CrossDesert(tt.input)
			assert.Equal(t, tt.expected, a)
		})
	}
}
