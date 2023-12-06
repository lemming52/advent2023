package d06waitforit

import (
	"advent/solutions/utils"
	"math"
	"strconv"
	"strings"
)

func solve(T, D float64) (int, int) {
	det := math.Sqrt(T*T - 4*D)
	return int(math.Floor((T-det)/2)) + 1, int(math.Ceil((T+det)/2)) - 1
}

func MarginError(lines []string) int {
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	output := 1
	for i, t := range times {
		min, max := solve(float64(utils.Stoi(t)), float64(utils.Stoi(distances[i])))
		output *= (max - min + 1)
	}
	return output
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(MarginError(lines)), "B"
}
