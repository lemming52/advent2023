package d15lenslibrary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAlgorithm(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "HASH",
			expected: 52,
		}, {
			input:    "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			expected: 1320,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.input), func(t *testing.T) {
			a := HashAlgorithm(tt.input)
			assert.Equal(t, tt.expected, a)
		})
	}
}
