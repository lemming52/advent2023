package d08hauntedwasteland

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrossDesert(t *testing.T) {
	tests := []struct {
		name      string
		expectedA int
		expectedB int
		input     []string
	}{
		{
			name:      "short",
			expectedA: 2,
			expectedB: 3,
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
			name:      "long",
			expectedA: 6,
			expectedB: 7,
			input: []string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
		}, {
			name:      "b",
			expectedA: 2,
			expectedB: 6,
			input: []string{
				"LR",
				"",
				"AAA = (11B, XXX)",
				"11B = (XXX, ZZZ)",
				"ZZZ = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"XXX = (XXX, XXX)",
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.name), func(t *testing.T) {
			a, b := CrossDesert(tt.input)
			assert.Equal(t, tt.expectedA, a)
			assert.Equal(t, tt.expectedB, b)
		})
	}
}
