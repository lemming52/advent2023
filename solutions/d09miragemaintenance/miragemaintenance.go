package d09miragemaintenance

import (
	"advent/solutions/utils"
	"strconv"
	"strings"
)

func PredictSequences(lines []string) (int, int) {
	low, high := 0, 0
	for _, l := range lines {
		l, h := predictSequence(l)
		low += l
		high += h
	}
	return high, low
}

func predictSequence(s string) (int, int) {
	vals := convertToInts(s)
	low, high := recursiveDifference(vals)
	return vals[0] - low, vals[len(vals)-1] + high
}

func convertToInts(s string) []int {
	components := strings.Split(s, " ")
	output := make([]int, len(components))
	for i, c := range components {
		output[i] = utils.Stoi(c)
	}
	return output
}

// recursiveDifference recursively calculates the differences of a slice of ints until all zero
// it then returns the new values for the preceding and subsequent predicted values of the slice
func recursiveDifference(vals []int) (int, int) {
	marker := len(vals) - 1
	new := make([]int, marker)
	allZero := true
	for i, v := range vals[:marker] {
		new[i] = vals[i+1] - v
		if new[i] != 0 {
			allZero = false
		}
	}
	if allZero {
		return 0, 0
	}
	low, high := recursiveDifference(new)
	return new[0] - low, new[marker-1] + high
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := PredictSequences(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
