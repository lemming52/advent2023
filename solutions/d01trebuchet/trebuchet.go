package d01trebuchet

import (
	"advent/solutions/utils"
	"strconv"
)

const CHAR_OFFSET = 48
const CHAR_MAX = 57

func findEmbeddedNumber(s string) int {
	var tens int
	for _, c := range s {
		if CHAR_OFFSET <= c && c <= CHAR_MAX {
			tens = int(c - CHAR_OFFSET)
			break
		}
	}
	for counter := len(s) - 1; counter >= 0; counter-- {
		if CHAR_OFFSET <= s[counter] && s[counter] <= CHAR_MAX {
			return tens*10 + int(s[counter]-CHAR_OFFSET)
		}
	}
	return 0
}

func SimpleCalibrationValues(lines []string) int {
	total := 0
	for _, l := range lines {
		total += findEmbeddedNumber(l)
	}
	return total
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(SimpleCalibrationValues(lines)), "B"
}
