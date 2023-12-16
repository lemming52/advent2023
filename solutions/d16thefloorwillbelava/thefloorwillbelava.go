package d16thefloorwillbelava

import (
	"advent/solutions/utils"
	"fmt"
	"strconv"
	"strings"
)

// Enum of cardinal directions
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type LaserGrid struct {
	grid           []string
	energisedCount int
	energyGrid     [][]bool
	xMax, yMax     int
	visited        [][]map[Direction]bool
}

func newLaserGrid(lines []string) *LaserGrid {
	grid := make([][]bool, len(lines))
	visited := make([][]map[Direction]bool, len(lines))
	for i, l := range lines {
		grid[i] = make([]bool, len(l))
		visited[i] = make([]map[Direction]bool, len(l))
		for j := range l {
			visited[i][j] = map[Direction]bool{North: false, East: false, South: false, West: false}
		}
	}
	return &LaserGrid{
		grid:           lines,
		energisedCount: 0,
		energyGrid:     grid,
		xMax:           len(lines[0]),
		yMax:           len(lines),
		visited:        visited,
	}
}

func (g *LaserGrid) isOnGrid(m *Move) bool {
	if m.x >= 0 && m.y >= 0 && m.x < g.xMax && m.y < g.yMax {
		return true
	}
	return false
}

func (g *LaserGrid) isDegenerate(m *Move) bool {
	visited := g.visited[m.y][m.x][m.entry]
	if !visited {
		g.visited[m.y][m.x][m.entry] = true
		return false
	}
	return true
}

func (g *LaserGrid) energise(x, y int) {
	if !g.energyGrid[y][x] {
		g.energisedCount += 1
	}
	g.energyGrid[y][x] = true
}

func (g *LaserGrid) print() {
	for _, r := range g.energyGrid {
		row := make([]string, len(r))
		for i, e := range r {
			if e {
				row[i] = "#"
			} else {
				row[i] = "."
			}
		}
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println("-----")
}

type Move struct {
	x, y  int
	entry Direction
}

func (m *Move) nextMove() *Move {
	switch m.entry {
	case North:
		m.y += 1
	case East:
		m.x -= 1
	case South:
		m.y -= 1
	case West:
		m.x += 1
	}
	return m
}

func (m *Move) shift(x, y int, d Direction) *Move {
	m.x += x
	m.y += y
	m.entry = d
	return m
}

func Energise(lines []string) int {
	grid := newLaserGrid(lines)
	firstMove := &Move{
		x:     0,
		y:     0,
		entry: West,
	}
	followLaser(firstMove, grid, 0)
	return grid.energisedCount
}

func followLaser(m *Move, grid *LaserGrid, i int) {
	if !grid.isOnGrid(m) || grid.isDegenerate(m) {
		return
	}
	grid.energise(m.x, m.y)
	c := grid.grid[m.y][m.x]
	if c == '.' {
		followLaser(m.nextMove(), grid, i+1)
		return
	}
	moves := nextMove(c, m)
	for _, m := range moves {
		followLaser(m, grid, i+1)
	}
}

func nextMove(c byte, m *Move) []*Move {
	char := rune(c)
	switch char {
	case '\\':
		return leftMirror(m)
	case '/':
		return rightMirror(m)
	case '-':
		if m.entry == East || m.entry == West {
			return []*Move{m.nextMove()}
		}
		return splitEastWest(m)
	case '|':
		if m.entry == North || m.entry == South {
			return []*Move{m.nextMove()}
		}
		return splitNorthSouth(m)
	default:
		return nil
	}
}

func leftMirror(m *Move) []*Move {
	switch m.entry {
	case North:
		return []*Move{m.shift(1, 0, West)}
	case East:
		return []*Move{m.shift(0, -1, South)}
	case South:
		return []*Move{m.shift(-1, 0, East)}
	case West:
		return []*Move{m.shift(0, 1, North)}
	}
	return nil
}

func rightMirror(m *Move) []*Move {
	switch m.entry {
	case North:
		return []*Move{m.shift(-1, 0, East)}
	case East:
		return []*Move{m.shift(0, 1, North)}
	case South:
		return []*Move{m.shift(1, 0, West)}
	case West:
		return []*Move{m.shift(0, -1, South)}
	}
	return nil
}

func splitEastWest(m *Move) []*Move {
	return []*Move{
		{x: m.x + 1, y: m.y, entry: West},
		{x: m.x - 1, y: m.y, entry: East},
	}
}

func splitNorthSouth(m *Move) []*Move {
	return []*Move{
		{x: m.x, y: m.y + 1, entry: North},
		{x: m.x, y: m.y - 1, entry: South},
	}
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a := Energise(lines)
	return strconv.Itoa(a), "B"
}
