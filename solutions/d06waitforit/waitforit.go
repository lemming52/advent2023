package d06waitforit

import (
	"advent/solutions/utils"
	"math"
	"strconv"
	"strings"
)

// solve t^2 - Tt + D = 0
func solve(T, D float64) (int, int) {
	det := math.Sqrt(T*T - 4*D)
	return int(math.Floor((T-det)/2)) + 1, int(math.Ceil((T+det)/2)) - 1
}

func MarginError(lines []string) (int, int) {
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	output := 1
	timeConcat, distanceConcat := 0, 0
	for i, t := range times {
		T, D := utils.Stoi(t), utils.Stoi(distances[i])
		min, max := solve(float64(T), float64(D))
		output *= (max - min + 1)
		timeConcat = timeConcat*int(math.Pow10(len(t))) + T
		distanceConcat = distanceConcat*int(math.Pow10(len(distances[i]))) + D
	}
	min, max := solve(float64(timeConcat), float64(distanceConcat))
	return output, max - min + 1
}

// This one just involved solving a quadratic equation.
func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := MarginError(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
