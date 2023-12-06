package d06waitforit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarginError(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	a, b := MarginError(input)
	assert.Equal(t, 288, a)
	assert.Equal(t, 71503, b)
}
