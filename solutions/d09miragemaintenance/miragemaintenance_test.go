package d09miragemaintenance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredictSequences(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	a, b := PredictSequences(input)
	assert.Equal(t, 114, a)
	assert.Equal(t, 2, b)
}
