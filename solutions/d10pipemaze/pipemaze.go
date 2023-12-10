package d10pipemaze

import (
	"advent/solutions/utils"
	"strconv"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func buildGrid(xMax, yMax int) [][]bool {
	grid := make([][]bool, yMax)
	for j := 0; j < yMax; j++ {
		grid[j] = make([]bool, xMax)
	}
	return grid
}

func findFirstMove(pipes []string, x, y int) (int, int, Direction) {
	c := pipes[y-1][x]
	if c == '|' || c == '7' || c == 'F' {
		return x, y - 1, North
	}
	c = pipes[y][x+1]
	if c == '-' || c == '7' || c == 'J' {
		return x + 1, y, East
	}
	c = pipes[y+1][x]
	if c == '|' || c == 'L' || c == 'J' {
		return x, y + 1, South
	}
	return x - 1, y, West
}

func TraversePipe(lines []string, startX, startY int) int {
	nextX, nextY, direction := findFirstMove(lines, startX, startY)
	nextC := lines[nextY][nextX]
	distance := 1
	for nextC != 'S' {
		nextX, nextY, direction = nextMove(nextC, nextX, nextY, direction)
		nextC = lines[nextY][nextX]
		distance += 1
	}
	return distance / 2
}

func nextMove(c byte, x, y int, d Direction) (int, int, Direction) {
	char := rune(c)
	switch char {
	case 'L':
		if d == South {
			return x + 1, y, East
		}
		return x, y - 1, North
	case 'J':
		if d == South {
			return x - 1, y, West
		}
		return x, y - 1, North
	case '7':
		if d == East {
			return x, y + 1, South
		}
		return x - 1, y, West
	case 'F':
		if d == West {
			return x, y + 1, South
		}
		return x + 1, y, East
	case '|':
		if d == South {
			return x, y + 1, South
		}
		return x, y - 1, North
	case '-':
		if d == East {
			return x + 1, y, East
		}
		return x - 1, y, West
	default:
		return x, y, North
	}
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a := TraversePipe(lines, 62, 90)
	return strconv.Itoa(a), "B"
}
