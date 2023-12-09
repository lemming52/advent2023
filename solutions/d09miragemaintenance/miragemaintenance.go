package d09miragemaintenance

import (
	"advent/solutions/utils"
	"strconv"
	"strings"
)

func PredictSequence(s string) (int, int) {
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

func PredictSequences(lines []string) (int, int) {
	low, high := 0, 0
	for _, l := range lines {
		l, h := PredictSequence(l)
		low += l
		high += h
	}
	return high, low
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := PredictSequences(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
