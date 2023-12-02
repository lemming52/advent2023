package d02cubeconundrum

import (
	"advent/solutions/utils"
	"strconv"
	"strings"
)

// const cubePattern = `(\d+) (red|blue|green)`

type Game struct {
	cubes map[string]int
}

func (g *Game) playGame(s string) bool {
	game := strings.Split(s, ":")
	hands := strings.Split(game[1], ";")
	for _, hand := range hands {
		if !g.possible(hand) {
			return false
		}
	}
	return true
}

func (g *Game) possible(s string) bool {
	sets := strings.Split(s, ",")
	for _, set := range sets {
		components := strings.Split(strings.Trim(set, " "), " ")
		count := utils.Stoi(components[0])
		if count > g.cubes[components[1]] {
			return false
		}
	}
	return true
}

func PlayGames(lines []string) int {
	game := Game{
		cubes: map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		},
	}
	total := 0
	for i, l := range lines {
		if game.playGame(l) {
			total += (i + 1)
		}
	}
	return total
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(PlayGames(lines)), "B"
}
