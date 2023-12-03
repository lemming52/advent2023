package d02cubeconundrum

import (
	"advent/solutions/utils"
	"strconv"
	"strings"
)

type Game struct {
	cubes      map[string]int
	minimumSet map[string]int
	possible   bool
}

func NewGame(r, g, b int) *Game {
	return &Game{
		cubes: map[string]int{
			"red":   r,
			"green": g,
			"blue":  b,
		},
		minimumSet: map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		},
		possible: true,
	}
}

func (g *Game) reset() {
	g.possible = true
	for k := range g.minimumSet {
		g.minimumSet[k] = 0
	}
}

func (g *Game) PlayGame(s string) (bool, int) {
	g.reset()
	game := strings.Split(s, ":")
	hands := strings.Split(game[1], ";")
	for _, hand := range hands {
		g.playHand(hand)
	}
	return g.possible, g.power()
}

func (g *Game) playHand(s string) {
	sets := strings.Split(s, ",")
	for _, set := range sets {
		components := strings.Split(strings.Trim(set, " "), " ")
		count := utils.Stoi(components[0])
		if count > g.minimumSet[components[1]] {
			g.minimumSet[components[1]] = count
		}
		if count > g.cubes[components[1]] {
			g.possible = false
		}
	}
}

func (g *Game) power() int {
	val := 1
	for _, v := range g.minimumSet {
		val *= v
	}
	return val
}

func PlayGames(lines []string) (int, int) {
	game := NewGame(12, 13, 14)
	total, totalPower := 0, 0
	for i, l := range lines {
		possible, power := game.PlayGame(l)
		if possible {
			total += (i + 1)
		}
		totalPower += power
	}
	return total, totalPower
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := PlayGames(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
