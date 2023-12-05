package d05seedfertiliser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSeeds(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
	a, b := MapSeeds(input)
	assert.Equal(t, 35, a)
	assert.Equal(t, 46, b)
}

func TestMapRanges(t *testing.T) {
	inputRanges := [][]int{
		{79, 79 + 13},
		{55, 55 + 12},
	}
	inputTransforms := []*Transform{
		{
			minRange: 98,
			maxRange: 99,
			offset:   50 - 98,
		}, {
			minRange: 50,
			maxRange: 97,
			offset:   52 - 50,
		},
	}
	fmt.Println(mapRanges(inputRanges, inputTransforms))
	assert.Equal(t, 0, 1)
}
