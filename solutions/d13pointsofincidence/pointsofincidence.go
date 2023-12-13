package d13pointsofincidence

import (
	"advent/solutions/utils"
	"strconv"
)

type Plate struct {
	grid         []string
	xMax, yMax   int
	xReflections []bool
	yReflections []bool
	optionCount  int
}

func newPlate(lines []string) *Plate {
	p := &Plate{
		grid:         lines,
		xMax:         len(lines[0]),
		yMax:         len(lines),
		xReflections: make([]bool, len(lines[0])-1),
		yReflections: make([]bool, len(lines)-1),
		optionCount:  0,
	}
	for i := range p.xReflections {
		p.xReflections[i] = true
		p.optionCount += 1
	}
	for i := range p.yReflections {
		p.yReflections[i] = true
		p.optionCount += 1
	}
	return p
}

func (p *Plate) checkReflections() int {
	for j := 0; j < p.yMax; j++ {
		for i := 0; i < p.xMax; i++ {
			for g, x := range p.xReflections {
				if !x {
					continue
				}
				ok, mX, mY := p.mirrorX(i, j, g)
				if !ok {
					continue
				}
				c := p.grid[j][i]
				if p.grid[mY][mX] != c {
					p.xReflections[g] = false
					p.optionCount -= 1
				}
				stop, isX, index := p.shouldStop()
				if stop {
					return p.evaluate(isX, index)
				}
			}
			for h, y := range p.yReflections {
				if !y {
					continue
				}
				ok, mX, mY := p.mirrorY(i, j, h)
				if !ok {
					continue
				}
				c := p.grid[j][i]
				if p.grid[mY][mX] != c {
					p.yReflections[h] = false
					p.optionCount -= 1
				}
				stop, isX, index := p.shouldStop()
				if stop {
					return p.evaluate(isX, index)
				}
			}
		}
	}
	return 0
}

func (p *Plate) mirrorX(x, y, index int) (bool, int, int) {
	newX := (index - x + 1) + index
	if !p.isValid(newX, y) {
		return false, 0, 0
	}
	return true, newX, y
}

func (p *Plate) mirrorY(x, y, index int) (bool, int, int) {
	newY := (index - y + 1) + index
	if !p.isValid(x, newY) {
		return false, 0, 0
	}
	return true, x, newY
}

func (p *Plate) isValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < p.xMax && y < p.yMax
}

func (p *Plate) shouldStop() (bool, bool, int) {
	if p.optionCount != 1 {
		return false, false, 0
	}
	for i, g := range p.xReflections {
		if g {
			return true, true, i
		}
	}
	for j, h := range p.yReflections {
		if h {
			return true, false, j
		}
	}
	return false, false, 0
}

func (p *Plate) evaluate(isX bool, index int) int {
	if isX {
		return index + 1
	}
	return (index + 1) * 100
}

func CheckPlates(lines []string) int {
	plates := assemblePlates(lines)
	total := 0
	for _, p := range plates {
		total += p.checkReflections()
	}
	return total
}

func assemblePlates(lines []string) []*Plate {
	plates := []*Plate{}
	buffer := []string{}
	for _, l := range lines {
		if l != "" {
			buffer = append(buffer, l)
			continue
		}
		plates = append(plates, newPlate(buffer))
		buffer = []string{}
	}
	if len(buffer) != 0 {
		plates = append(plates, newPlate(buffer))
	}
	return plates
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a := CheckPlates(lines)
	return strconv.Itoa(a), "B"
}
