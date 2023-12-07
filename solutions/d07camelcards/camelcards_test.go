package d07camelcards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreHands(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	a := ScoreHands(input)
	assert.Equal(t, 6440, a)
}
