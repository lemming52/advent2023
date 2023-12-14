package d14parabolicreflectordish

import (
	"advent/solutions/utils"
	"fmt"
	"strconv"
	"strings"
)

type Rock int

const (
	Space Rock = iota
	Round
	Cube
)

type Dish struct {
	grid [][]Rock
	yMax int
}

func (d *Dish) Tilt() {
	for i := range d.grid[0] {
		d.tiltColumn(i)
	}
}

func (d *Dish) tiltColumn(col int) {
	currentRest := 0
	for j, r := range d.grid {
		rock := r[col]
		switch rock {
		case Round:
			r[col] = Space
			d.grid[currentRest][col] = rock
			currentRest += 1
		case Cube:
			currentRest = j + 1
		default:
			continue
		}
	}
}

func (d *Dish) print() {
	for _, r := range d.grid {
		row := make([]string, len(r))
		for i, rock := range r {
			switch rock {
			case Round:
				row[i] = "O"
			case Cube:
				row[i] = "#"
			default:
				row[i] = "."
			}
		}
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println("-----")
}

func (d *Dish) CalculateLoad() int {
	total := 0
	for j, r := range d.grid {
		for _, rock := range r {
			if rock == Round {
				total += d.yMax - j
			}
		}
	}
	return total
}

func FocusDish(lines []string) int {
	d := buildDish(lines)
	d.Tilt()
	return d.CalculateLoad()
}

func buildDish(lines []string) *Dish {
	d := &Dish{grid: make([][]Rock, len(lines)), yMax: len(lines)}
	for j, l := range lines {
		d.grid[j] = make([]Rock, len(l))
		for i, c := range l {
			switch c {
			case 'O':
				d.grid[j][i] = Round
			case '#':
				d.grid[j][i] = Cube
			default:
				continue
			}

		}
	}
	return d
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a := FocusDish(lines)
	return strconv.Itoa(a), "B"
}
