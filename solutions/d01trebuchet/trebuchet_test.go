package d01trebuchet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCalibrationValues(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	res := SimpleCalibrationValues(input)
	assert.Equal(t, 142, res)
}

func TestWordCalibrationValues(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	res := WordCalibrationValues(input)
	assert.Equal(t, 281, res)
}
