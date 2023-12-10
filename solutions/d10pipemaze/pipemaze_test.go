package d10pipemaze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraversePipeA(t *testing.T) {
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
			a, _ := TraversePipe(tt.input, tt.x, tt.y)
			assert.Equal(t, tt.expected, a)
		})
	}
}

func TestTraversePipeB(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
		x, y     int
	}{
		{
			name: "simple",
			input: []string{
				"...........",
				".S-------7.",
				".|F-----7|.",
				".||.....||.",
				".||.....||.",
				".|L-7.F-J|.",
				".|..|.|..|.",
				".L--J.L--J.",
				"...........",
			},
			expected: 4,
			x:        1,
			y:        1,
		}, {
			name: "longer",
			input: []string{
				"..........",
				".S------7.",
				".|F----7|.",
				".||OOOO||.",
				".||OOOO||.",
				".|L-7F-J|.",
				".|II||II|.",
				".L--JL--J.",
				"..........",
			},
			expected: 4,
			x:        1,
			y:        1,
		}, {
			name: "bigger",
			input: []string{
				"OF----7F7F7F7F-7OOOO",
				"O|F--7||||||||FJOOOO",
				"O||OFJ||||||||L7OOOO",
				"FJL7L7LJLJ||LJIL-7OO",
				"L--JOL7IIILJS7F-7L7O",
				"OOOOF-JIIF7FJ|L7L7L7",
				"OOOOL7IF7||L7|IL7L7|",
				"OOOOO|FJLJ|FJ|F7|OLJ",
				"OOOOFJL-7O||O||||OOO",
				"OOOOL---JOLJOLJLJOOO",
			},
			expected: 8,
			x:        12,
			y:        4,
		}, {
			name: "biggerstill",
			input: []string{
				"FF7FSF7F7F7F7F7F---7",
				"L|LJ||||||||||||F--J",
				"FL-7LJLJ||||||LJL-77",
				"F--JF--7||LJLJ7F7FJ-",
				"L---JF-JLJ.||-FJLJJ7",
				"|F|F-JF---7F7-L7L|7|",
				"|FFJF7L7F-JF7|JL---7",
				"7-L-JL7||F7|L7F-7F7|",
				"L.L7LFJ|||||FJL7||LJ",
				"L7JLJL-JLJLJL--JLJ.L",
			},
			expected: 10,
			x:        4,
			y:        0,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.name), func(t *testing.T) {
			_, a := TraversePipe(tt.input, tt.x, tt.y)
			assert.Equal(t, tt.expected, a)
		})
	}
}
