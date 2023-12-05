package d05seedfertiliser

import (
	"advent/solutions/utils"
	"math"
	"strconv"
	"strings"
)

type Transform struct {
	minRange, maxRange int
	offset             int
}

func (t *Transform) MapValue(v int) (int, bool) {
	if t.minRange <= v && v <= t.maxRange {
		return v + t.offset, true
	}
	return 0, false
}

func MapSeeds(lines []string) int {
	seeds := parseSeeds(lines[0])
	transforms := [][]*Transform{}
	index := 3
	ranges := []string{}
	for index < len(lines) {
		l := lines[index]
		if l == "" {
			transforms = append(transforms, buildTransformStep(ranges))
			ranges = []string{}
			index += 2
			continue
		}
		ranges = append(ranges, l)
		index += 1
	}
	if len(ranges) != 0 {
		transforms = append(transforms, buildTransformStep(ranges))
	}
	minimum := math.MaxInt
	for _, s := range seeds {
		for _, set := range transforms {
			s = checkTransforms(s, set)
		}
		if s < minimum {
			minimum = s
		}
	}
	return minimum
}

func checkTransforms(val int, transforms []*Transform) int {
	for _, t := range transforms {
		v, ok := t.MapValue(val)
		if ok {
			return v
		}
	}
	return val
}

func parseSeeds(s string) []int {
	components := strings.Split(s, ":")
	seeds := strings.Split(strings.Trim(components[1], " "), " ")
	output := make([]int, len(seeds))
	for i, s := range seeds {
		output[i] = utils.Stoi(s)
	}
	return output
}

func buildTransformStep(lines []string) []*Transform {
	ranges := make([]*Transform, len(lines))
	for i, l := range lines {
		components := strings.Split(l, " ")
		dest, source, window := utils.Stoi(components[0]), utils.Stoi(components[1]), utils.Stoi(components[2])
		ranges[i] = &Transform{
			minRange: source,
			maxRange: source + window - 1,
			offset:   dest - source,
		}
	}
	return ranges
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(MapSeeds(lines)), "B"
}
