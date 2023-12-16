package d15lenslibrary

import (
	"advent/solutions/utils"
	"strconv"
	"strings"
)

func HashAlgorithm(line string) int {
	v := 0
	for _, l := range strings.Split(line, ",") {
		v += hashString(l)
	}
	return v
}

func hashString(s string) int {
	v := 0
	for _, c := range s {
		v = hashChar(c, v)
	}
	return v
}

func hashChar(c rune, v int) int {
	v += int(c)
	v *= 17
	v = v % 256
	return v
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a := HashAlgorithm(lines[0])
	return strconv.Itoa(a), "B"
}
