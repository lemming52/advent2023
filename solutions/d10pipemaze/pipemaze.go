package d10pipemaze

import (
	"advent/solutions/utils"
	"strconv"
)

// Enum of cardinal directions
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// buildGrid builds an empty grid of bools to size
func buildGrid(xMax, yMax int) [][]bool {
	grid := make([][]bool, yMax)
	for j := 0; j < yMax; j++ {
		grid[j] = make([]bool, xMax)
	}
	return grid
}

// findFirstMove looks at the start position to find one of the connected pipes
// and the direction that pipe is moved into
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

// TraversePipe moves all across the pipe, totalling the distance
// as the pipe is traversed pockets are spotted to be on either
// the right or left of the pipe
// finally the pockets of not pipe are explored and area counted
func TraversePipe(lines []string, startX, startY int) (int, int) {
	xMax, yMax := len(lines[0]), len(lines)
	visited, left := buildGrid(xMax, yMax), buildGrid(xMax, yMax)
	nextX, nextY, firstDirection := findFirstMove(lines, startX, startY)
	inDirection := firstDirection
	visited[nextY][nextX] = true
	left[nextY][nextX] = false
	nextC := lines[nextY][nextX]
	distance := 1
	for nextC != 'S' {
		nX, nY, outDirection := nextMove(nextC, nextX, nextY, inDirection)
		handleOthers(nextX, nextY, xMax, yMax, inDirection, outDirection, visited, left)
		nextX, nextY = nX, nY
		inDirection = outDirection
		nextC = lines[nextY][nextX]
		visited[nextY][nextX] = true
		left[nextY][nextX] = false
		distance += 1
	}
	handleOthers(startX, startY, xMax, yMax, inDirection, firstDirection, visited, left)
	return distance / 2, countPockets(xMax, yMax, left, visited)
}

// nextMove evaluates a pipe segment, returning the next coordinate and
// the direction that segment is entered
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

// handleOthers flags adjacent cells to the left of the pipe
func handleOthers(x, y, xMax, yMax int, in, out Direction, visited, left [][]bool) {
	leftN := selectLeftUnvisited(x, y, in, out)
	for _, xy := range leftN {
		if xy[0] >= 0 && xy[1] >= 0 && xy[0] < xMax && xy[1] < yMax {
			if !visited[xy[1]][xy[0]] {
				left[xy[1]][xy[0]] = true
			}
		}
	}
}

// selectLeftUnvisited returns the adjacent unvisited neighbours for any in/out traversal of a cell
func selectLeftUnvisited(x, y int, in, out Direction) [][]int {
	if in == South {
		switch out {
		case East:
			return nil
		case South:
			return [][]int{{x + 1, y}}
		case West:
			return [][]int{{x + 1, y}, {x, y + 1}}
		}
	}
	if in == West {
		switch out {
		case South:
			return nil
		case West:
			return [][]int{{x, y + 1}}
		case North:
			return [][]int{{x - 1, y}, {x, y + 1}}
		}
	}
	if in == North {
		switch out {
		case West:
			return nil
		case North:
			return [][]int{{x - 1, y}}
		case East:
			return [][]int{{x - 1, y}, {x, y - 1}}
		}
	}
	if in == East {
		switch out {
		case North:
			return nil
		case East:
			return [][]int{{x, y - 1}}
		case South:
			return [][]int{{x + 1, y}, {x, y - 1}}
		}
	}
	return nil
}

// countPockets recursively explores pockets of not pipe in the space, and totals
// the area to the left and right of the pipe. returns whichever is smaller
// I know that's an assumption, but it worked.
func countPockets(xMax, yMax int, left, visited [][]bool) int {
	leftTotal, rightTotal := 0, 0
	explored := buildGrid(xMax, yMax)
	xMax, yMax = xMax-1, yMax-1
	for j := 0; j <= yMax; j++ {
		for i := 0; i <= xMax; i++ {
			if visited[j][i] || explored[j][i] {
				continue
			}
			t, l := exploreCell(i, j, xMax, yMax, left, visited, explored)
			if l {
				leftTotal += t
			} else {
				rightTotal += t
			}
		}
	}
	if leftTotal < rightTotal {
		return leftTotal
	}
	return rightTotal
}

// exploreCell recursively expands a pocket to its full extent, and computes the area
// also notes if the pocket is to the left or right of the pipe
func exploreCell(x, y, xMax, yMax int, left, visited, explored [][]bool) (int, bool) {
	neighbours := utils.All2DNeighbours(x, y, xMax, yMax)
	isLeft := left[y][x]
	total := 1
	explored[y][x] = true
	for _, xy := range neighbours {
		if visited[xy[1]][xy[0]] || explored[xy[1]][xy[0]] {
			continue
		}
		inc, l := exploreCell(xy[0], xy[1], xMax, yMax, left, visited, explored)
		total += inc
		isLeft = isLeft || l
	}
	return total, isLeft
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := TraversePipe(lines, 62, 90)
	return strconv.Itoa(a), strconv.Itoa(b)
}
