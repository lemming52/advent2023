package d10pipemaze

import (
	"advent/solutions/utils"
	"fmt"
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
	if y != 0 {
		c := pipes[y-1][x]
		if c == '|' || c == '7' || c == 'F' {
			return x, y - 1, North
		}
	}
	c := pipes[y][x+1]
	if c == '-' || c == '7' || c == 'J' {
		return x + 1, y, East
	}
	c = pipes[y+1][x]
	if c == '|' || c == 'L' || c == 'J' {
		return x, y + 1, South
	}
	return x - 1, y, West
}

func TraversePipe(lines []string, startX, startY int) (int, int) {
	xMax, yMax := len(lines[0]), len(lines)
	visited, left, right := buildGrid(xMax, yMax), buildGrid(xMax, yMax), buildGrid(xMax, yMax)
	nextX, nextY, firstDirection := findFirstMove(lines, startX, startY)
	inDirection := firstDirection
	visited[nextY][nextX] = true
	left[nextY][nextX], right[nextY][nextX] = false, false
	nextC := lines[nextY][nextX]
	distance := 1
	for nextC != 'S' {
		nX, nY, outDirection := nextMove(nextC, nextX, nextY, inDirection)
		handleOthers(nextX, nextY, xMax, yMax, inDirection, outDirection, visited, left, right)
		nextX, nextY = nX, nY
		inDirection = outDirection
		nextC = lines[nextY][nextX]
		visited[nextY][nextX] = true
		left[nextY][nextX], right[nextY][nextX] = false, false
		distance += 1
	}
	handleOthers(startX, startY, xMax, yMax, inDirection, firstDirection, visited, left, right)
	l, r := countArea(left), countArea(right)
	fmt.Println(l, r)
	if l < r {
		return distance / 2, l
	}
	return distance / 2, r
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

func handleOthers(x, y, xMax, yMax int, in, out Direction, visited, left, right [][]bool) {
	leftN, rightN := selectUnvisited(x, y, in, out)
	for _, xy := range leftN {
		if xy[0] >= 0 && xy[1] >= 0 && xy[0] < xMax && xy[1] < yMax {
			if !visited[xy[1]][xy[0]] {
				left[xy[1]][xy[0]] = true
			}
		}
	}
	for _, xy := range rightN {
		if xy[0] >= 0 && xy[1] >= 0 && xy[0] < xMax && xy[1] < yMax {
			if !visited[xy[1]][xy[0]] {
				right[xy[1]][xy[0]] = true
			}
		}
	}
}

func selectUnvisited(x, y int, in, out Direction) ([][]int, [][]int) {
	if in == South {
		switch out {
		case East:
			return nil, [][]int{{x - 1, y}, {x, y + 1}}
		case South:
			return [][]int{{x + 1, y}}, [][]int{{x - 1, y}}
		case West:
			return [][]int{{x + 1, y}, {x, y + 1}}, nil
		}
	}
	if in == West {
		switch out {
		case South:
			return nil, [][]int{{x - 1, y}, {x, y - 1}}
		case West:
			return [][]int{{x, y + 1}}, [][]int{{x, y - 1}}
		case North:
			return [][]int{{x - 1, y}, {x, y + 1}}, nil
		}
	}
	if in == North {
		switch out {
		case West:
			return nil, [][]int{{x + 1, y}, {x, y - 1}}
		case North:
			return [][]int{{x - 1, y}}, [][]int{{x + 1, y}}
		case East:
			return [][]int{{x - 1, y}, {x, y - 1}}, nil
		}
	}
	if in == East {
		switch out {
		case North:
			return nil, [][]int{{x + 1, y}, {x, y + 1}}
		case East:
			return [][]int{{x, y - 1}}, [][]int{{x, y + 1}}
		case South:
			return [][]int{{x + 1, y}, {x, y - 1}}, nil
		}
	}
	return nil, nil
}

func countArea(grid [][]bool) int {
	total := 0
	for _, y := range grid {
		for _, x := range y {
			if x {
				total += 1
			}
		}
	}
	return total
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := TraversePipe(lines, 62, 90)
	return strconv.Itoa(a), strconv.Itoa(b)
}
