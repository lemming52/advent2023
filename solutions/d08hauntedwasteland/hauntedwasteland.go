package d08hauntedwasteland

import (
	"advent/solutions/utils"
	"fmt"
	"regexp"
	"strconv"
)

const start = "AAA"
const end = "ZZZ"
const connection = `(\w{3}) = \((\w{3}), (\w{3})\)`

func CrossDesert(lines []string) (int, int) {
	instructions := lines[0]

	pattern := regexp.MustCompile(connection)
	connections := map[string][]string{}
	starts := []string{}
	for _, l := range lines[2:] {
		components := pattern.FindStringSubmatch(l)
		connections[components[1]] = []string{components[2], components[3]}
		if components[1][2] == 'A' {
			starts = append(starts, components[1])
		}
	}

	return SingleTrack(instructions, connections), MultiTrack(starts, instructions, connections)
}

func SingleTrack(instructions string, connections map[string][]string) int {
	i, steps, length := 0, 0, len(instructions)
	current := start
	var next string
	for true {
		if i == length {
			i = 0
		}
		steps += 1
		if instructions[i] == 'L' {
			next = connections[current][0]
		} else {
			next = connections[current][1]
		}
		if next == end {
			return steps
		}
		current = next
		i += 1
	}
	return 0
}

type CycleReport struct {
	position    string
	startOffset int
	length      int
	finalStates []int
}

func (c *CycleReport) isInEndState(steps int) (bool, int) {
	cyclePosition := (steps - c.startOffset) % c.length
	isEndState := false
	nextPosition := c.length
	for _, f := range c.finalStates {
		if f < cyclePosition {
			continue
		}
		if f == cyclePosition {
			isEndState = true
			continue
		}
		nextPosition = steps + f - cyclePosition
		break
	}
	return isEndState, nextPosition
}

func (c *CycleReport) print() {
	fmt.Println(c.position, c.startOffset, c.length, c.finalStates)
}

type EndState struct {
	name  string
	steps int
}

func MultiTrack(positions []string, instructions string, connections map[string][]string) int {
	cycles := make([]*CycleReport, len(positions))
	for j, p := range positions {
		cycle := &CycleReport{
			position:    p,
			finalStates: []int{},
		}
		i, steps, length := 0, 0, len(instructions)
		visited := map[string]int{}
		name := fmt.Sprintf("%s:%d", p, i)
		visited[name] = steps
		endStates := []*EndState{}
		for true {
			if i == length {
				i = 0
			}
			steps += 1
			if instructions[i] == 'L' {
				positions[j] = connections[positions[j]][0]
			} else {
				positions[j] = connections[positions[j]][1]
			}
			if isEndState(positions[j]) {
				endStates = append(endStates, &EndState{name: positions[j], steps: steps})
			}
			i += 1
			name := fmt.Sprintf("%s:%d", positions[j], i)
			v, ok := visited[name]
			if ok {
				cycle.startOffset = v
				cycle.length = steps - v
				break
			}
			visited[name] = steps
		}
		for _, e := range endStates {
			if e.steps < cycle.startOffset {
				continue
			}
			cycle.finalStates = append(cycle.finalStates, e.steps-cycle.startOffset)
		}
		cycles[j] = cycle
	}
	position := 0
	for _, c := range cycles {
		if c.startOffset > position {
			position = c.startOffset
		}
	}
	success := false
	for !success {
		nextOption := position
		success = true
		for _, c := range cycles {
			ok, v := c.isInEndState(position)
			if !ok {
				success = false
			}
			if v > nextOption {
				nextOption = v
			}
		}
		position = nextOption
	}
	return position
	/*
		i, steps, length := 0, 0, len(instructions)
		for true {
			if i == length {
				i = 0
			}
			steps += 1
			for j, p := range positions {
				if instructions[i] == 'L' {
					positions[j] = connections[p][0]
				} else {
					positions[j] = connections[p][1]
				}
			}
			if isFinalState(positions) {
				return steps
			}
			i += 1
		}
		return 0
	*/
}

func isEndState(s string) bool {
	return s[2] == 'Z'
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := CrossDesert(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
