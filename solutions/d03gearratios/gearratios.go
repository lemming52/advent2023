package d03gearratios

import (
	"advent/solutions/utils"
	"fmt"
	"strconv"
)

const runeOffset = 48

type Part struct {
	value int
	xMin  int
	xMax  int
	y     int
}

func newPart(val, x1, x2, y int) *Part {
	return &Part{
		value: val,
		xMin:  x1,
		xMax:  x2,
		y:     y,
	}
}

type Gear struct {
	value int
	count int
}

type Engine struct {
	Parts   []*Part
	Lines   []string
	Symbols [][]bool
	Gears   map[string]*Gear
	xMax    int
	yMax    int
}

func newEngine(x, y int) *Engine {
	e := &Engine{
		Parts: []*Part{},
		Gears: map[string]*Gear{},
		xMax:  x,
		yMax:  y,
	}
	e.Symbols = make([][]bool, y+1)
	for i := 0; i <= y; i++ {
		e.Symbols[i] = make([]bool, x+1)
	}
	return e
}

func (e *Engine) GetPartNeighbours(p *Part) [][]int {
	coords := [][]int{
		{p.xMin - 1, p.y - 1},
		{p.xMin - 1, p.y},
		{p.xMin - 1, p.y + 1},
		{p.xMax + 1, p.y - 1},
		{p.xMax + 1, p.y},
		{p.xMax + 1, p.y + 1},
	}
	for i := p.xMin; i <= p.xMax; i++ {
		coords = append(coords, [][]int{{i, p.y - 1}, {i, p.y + 1}}...)
	}
	correct := [][]int{}
	for _, xy := range coords {
		if xy[0] >= 0 && xy[1] >= 0 && xy[0] <= e.xMax && xy[1] <= e.yMax {
			correct = append(correct, xy)
		}
	}
	return correct
}

func (e *Engine) TotalPartNumbers() int {
	total := 0
	for _, p := range e.Parts {
		total += e.CheckPart(p)
	}
	return total
}

func (e *Engine) CheckPart(p *Part) int {
	value := 0
	neighbours := e.GetPartNeighbours(p)
	for _, xy := range neighbours {
		if e.Symbols[xy[1]][xy[0]] {
			value = p.value
			if e.Lines[xy[1]][xy[0]] == '*' {
				e.AddGear(xy[0], xy[1], p)
			}
		}
	}
	return value
}

func (e *Engine) AddGear(x, y int, p *Part) {
	v, ok := e.Gears[fmt.Sprintf("%d,%d", x, y)]
	if ok {
		v.count += 1
		v.value *= p.value
	} else {
		e.Gears[fmt.Sprintf("%d,%d", x, y)] = &Gear{
			value: p.value,
			count: 1,
		}
	}
}

func (e *Engine) CountGears() int {
	total := 0
	for _, v := range e.Gears {
		if v.count == 2 {
			total += v.value
		}
	}
	return total
}

func BuildEngine(lines []string) *Engine {
	e := newEngine(len(lines[0])-1, len(lines)-1)
	e.Lines = lines
	for y, l := range lines {
		val, magnitude := 0, 0
		for x, c := range l {
			switch c {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				val = val*10 + int(c-runeOffset)
				magnitude += 1
				continue
			case '.':
				if val != 0 {
					e.Parts = append(e.Parts, newPart(val, x-magnitude, x-1, y))
					val, magnitude = 0, 0
					continue
				}
			default:
				if val != 0 {
					e.Parts = append(e.Parts, newPart(val, x-magnitude, x-1, y))
					val, magnitude = 0, 0
				}
				e.Symbols[y][x] = true
			}
		}
		if val != 0 {
			e.Parts = append(e.Parts, newPart(val, e.xMax-magnitude, e.xMax, y))
		}
	}
	return e
}

func PartNumbers(lines []string) (int, int) {
	e := BuildEngine(lines)
	return e.TotalPartNumbers(), e.CountGears()
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := PartNumbers(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
