package d08hauntedwasteland

import (
	"advent/solutions/utils"
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

func MultiTrack(positions []string, instructions string, connections map[string][]string) int {
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
}

func isFinalState(positions []string) bool {
	for _, p := range positions {
		if p[2] != 'Z' {
			return false
		}
	}
	return true
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := CrossDesert(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
