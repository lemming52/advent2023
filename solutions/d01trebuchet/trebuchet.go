package d01trebuchet

import (
	"advent/solutions/utils"
	"regexp"
	"strconv"
)

const runeOffset = 48
const runeMax = 57
const digitPattern = `\d`
const numberPattern = `one|two|three|four|five|six|seven|eight|nine`

func SimpleCalibrationValues(lines []string) int {
	total := 0
	for _, l := range lines {
		total += findEmbeddedNumber(l)
	}
	return total
}

func findEmbeddedNumber(s string) int {
	var tens int
	for _, c := range s {
		if runeOffset <= c && c <= runeMax {
			tens = int(c - runeOffset)
			break
		}
	}
	for counter := len(s) - 1; counter >= 0; counter-- {
		if runeOffset <= s[counter] && s[counter] <= runeMax {
			return tens*10 + int(s[counter]-runeOffset)
		}
	}
	return 0
}

type Number struct {
	index int
	value int
}

func newNumber(s string, index int) *Number {
	return &Number{
		index: index,
		value: int(s[index] - runeOffset),
	}
}

func newWordNumber(s string, start, end, offset int) *Number {
	substring := s[start:end]
	var val int
	switch substring {
	case "one":
		val = 1
	case "two":
		val = 2
	case "three":
		val = 3
	case "four":
		val = 4
	case "five":
		val = 5
	case "six":
		val = 6
	case "seven":
		val = 7
	case "eight":
		val = 8
	case "nine":
		val = 9
	}
	return &Number{
		index: start + offset,
		value: val,
	}
}

func WordCalibrationValues(lines []string) int {
	total := 0
	digit := regexp.MustCompile(digitPattern)
	number := regexp.MustCompile(numberPattern)
	for _, l := range lines {
		total += findEmbeddedNumberRegex(l, digit, number)
	}
	return total
}

func findEmbeddedNumberRegex(s string, digit, number *regexp.Regexp) int {
	d1, d2 := convertDigitMatches(s, digit)
	n1, n2 := convertNumberMatches(s, number)
	if d1 == nil {
		return n1.value*10 + n2.value
	}
	if n1 == nil {
		return d1.value*10 + d2.value
	}
	var a, b *Number
	if d1.index < n1.index {
		a = d1
	} else {
		a = n1
	}
	if d2.index > n2.index {
		b = d2
	} else {
		b = n2
	}
	return a.value*10 + b.value
}

func convertDigitMatches(s string, digit *regexp.Regexp) (*Number, *Number) {
	matches := digit.FindAllStringSubmatchIndex(s, -1)
	if len(matches) == 0 {
		return nil, nil
	}
	first := matches[0]
	last := matches[len(matches)-1]
	return newNumber(s, first[0]), newNumber(s, last[0])
}

func convertNumberMatches(s string, number *regexp.Regexp) (*Number, *Number) {
	matches := number.FindAllStringSubmatchIndex(s, -1)
	if len(matches) == 0 {
		return nil, nil
	}
	first := matches[0]
	return newWordNumber(s, first[0], first[1], 0), guaranteeLast(s, 0, matches[len(matches)-1], number)
}

func guaranteeLast(s string, offset int, match []int, number *regexp.Regexp) *Number {
	furtherMatches := number.FindAllStringSubmatchIndex(s[match[0]+1:], -1)
	if furtherMatches == nil {
		return newWordNumber(s, match[0], match[1], offset)
	}
	return guaranteeLast(s[match[0]+1:], offset+match[0]+1, furtherMatches[0], number)

}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(SimpleCalibrationValues(lines)), strconv.Itoa(WordCalibrationValues(lines))
}
