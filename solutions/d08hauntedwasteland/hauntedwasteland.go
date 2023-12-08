package d08hauntedwasteland

import (
	"advent/solutions/utils"
	"regexp"
	"strconv"
)

const start = "AAA"
const end = "ZZZ"
const connection = `(\w{3}) = \((\w{3}), (\w{3})\)`

func CrossDesert(lines []string) int {
	instructions := lines[0]

	pattern := regexp.MustCompile(connection)
	connections := map[string][]string{}
	for _, l := range lines[2:] {
		components := pattern.FindStringSubmatch(l)
		connections[components[1]] = []string{components[2], components[3]}
	}

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

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a := CrossDesert(lines)
	return strconv.Itoa(a), "B"
}
