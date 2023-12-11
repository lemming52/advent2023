package d05seedfertiliser

import (
	"advent/solutions/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Transform encapsulates a mapping range, the start and end values and the offset to add
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

func MapSeeds(lines []string) (int, int) {
	transforms := buildTransforms(lines)
	seeds := parseSeeds(lines[0])
	singleSeed := singleSeeds(seeds, transforms)
	return singleSeed, seedRanges(seeds, transforms)
}

// buildTransforms computes the transform ranges for each step in the mapping
func buildTransforms(lines []string) [][]*Transform {
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
	return transforms
}

// buildTransformStep computes all transform ranges for a single map step
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

func parseSeeds(s string) []int {
	components := strings.Split(s, ":")
	seeds := strings.Split(strings.Trim(components[1], " "), " ")
	output := make([]int, len(seeds))
	for i, s := range seeds {
		output[i] = utils.Stoi(s)
	}
	return output
}

func convertToRanges(seeds []int) [][]int {
	seedRanges := make([][]int, len(seeds)/2)
	for i := 0; i < len(seeds)/2; i += 1 {
		seedRanges[i] = []int{seeds[2*i], seeds[2*i] + seeds[2*i+1]}
	}
	return seedRanges
}

// singleSeeds solves for when the transforms are maps for single values
func singleSeeds(seeds []int, transforms [][]*Transform) int {
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

// checkTransform transforms a given value for a single mapping step
func checkTransforms(val int, transforms []*Transform) int {
	for _, t := range transforms {
		v, ok := t.MapValue(val)
		if ok {
			return v
		}
	}
	return val
}

// mapRanges changes values based on the transforms
// given a range of values, it applies a transform to that range where applicable
// and splits off the unapplicable sections to try again with the other transforms
// it's a bit awkward
func mapRanges(ranges [][]int, transforms []*Transform) [][]int {
	newRanges := [][]int{}
	for _, r := range ranges {
		minVal, maxVal := r[0], r[1]
		for _, t := range transforms {
			if minVal == maxVal {
				break
			}
			if t.minRange <= minVal {
				if t.maxRange <= minVal {
					// no overlap above
				} else if maxVal <= t.maxRange {
					// complete contain
					newRanges = append(newRanges, []int{minVal + t.offset, maxVal + t.offset})
					minVal, maxVal = 0, 0
				} else {
					// lower overlap
					newRanges = append(newRanges, []int{minVal + t.offset, t.maxRange + t.offset})
					minVal = t.maxRange + 1
				}
			} else {
				if maxVal <= t.minRange {
					// no overlap below
				} else if maxVal <= t.maxRange {
					// upper overlap
					newRanges = append(newRanges, []int{t.minRange + t.offset, maxVal + t.offset})
					maxVal = t.minRange - 1
				} else {
					// range contains transform
					newRanges = append(newRanges, []int{t.minRange + t.offset, t.maxRange + t.offset})
					ranges = append(ranges, []int{t.maxRange + 1, maxVal})
					maxVal = t.minRange - 1
				}
			}
		}
		if minVal != maxVal {
			newRanges = append(newRanges, []int{minVal, maxVal})
		}
	}
	sort.Slice(newRanges, func(i, j int) bool {
		return newRanges[i][0] < newRanges[j][0]
	})
	outputRanges := [][]int{}
	i, j := 0, 1
	currentTop := newRanges[i][1]
	for i <= len(newRanges)-1 {
		if (i + j) >= len(newRanges) {
			outputRanges = append(outputRanges, []int{newRanges[i][0], newRanges[len(newRanges)-1][1]})
			break
		}
		if newRanges[i+j][0] <= newRanges[i][1] {
			currentTop = newRanges[i+j][1]
			j += 1
			continue
		}
		outputRanges = append(outputRanges, []int{newRanges[i][0], currentTop})
		currentTop = newRanges[i+j][1]
		i += j
		j = 1
	}
	return outputRanges
}

// seedRanges solves with seed ranges and returns the lowest
func seedRanges(seeds []int, transforms [][]*Transform) int {
	ranges := convertToRanges(seeds)
	for _, t := range transforms {
		ranges = mapRanges(ranges, t)
	}
	return ranges[0][0]
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := MapSeeds(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
