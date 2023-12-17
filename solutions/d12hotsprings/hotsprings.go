package d12hotsprings

import (
	"advent/solutions/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type BrokenSprings struct {
	count    int
	min, max int
	index    int
	starts   []int
	dof      int
}

func (b *BrokenSprings) print() {
	fmt.Printf("size: %d min: %d max: %d starts: %v dof: %d\n", b.count, b.min, b.max, b.starts, b.dof)
}

func (b *BrokenSprings) isPossibility(i int) bool {
	if len(b.starts) > 0 {
		return utils.Contains(i, b.starts)
	}
	return b.min <= i && i <= b.max
}

func ArrangeSprings(lines []string) int {
	total := 0
	for _, l := range lines {
		chars, nums := extractComponents(l)
		res := solve(chars, nums, map[string]int{})
		fmt.Println(l, res, "OUTPUT")
		total += res
	}
	return total
}

func buildModels(s string) int {
	chars, nums := extractComponents(s)
	fmt.Println(chars, nums)
	springs := defineSpringRanges(nums, len(chars))
	for i := 0; i < 5; i++ {
		chars = recursive(chars, springs)
		springs = defineSpringRanges(nums, len(chars))
	}
	return 0
}

func extractComponents(s string) (string, []int) {
	components := strings.Split(s, " ")
	sections := strings.Split(components[0], ".")
	out := []string{}
	for _, s := range sections {
		if s != "" {
			out = append(out, s)
		}
	}
	vals := strings.Split(components[1], ",")
	output := make([]int, len(vals))
	for i, v := range vals {
		output[i] = utils.Stoi(v)
	}
	return strings.Join(out, "."), output
}

func breakIntoSections(s string) ([][]int, [][]int) {
	sections, confirmed := [][]int{}, [][]int{}
	sectionLength, confirmedLength := 0, 0
	sI, cI := -1, -1
	for i, c := range s {
		switch c {
		case '#':
			if cI == -1 {
				cI = i
			}
			if sI == -1 {
				sI = i
			}
			confirmedLength += 1
			sectionLength += 1
		case '?':
			if sI == -1 {
				sI = i
			}
			if confirmedLength != 0 {
				confirmed = append(confirmed, []int{confirmedLength, cI})
				cI = -1
				confirmedLength = 0
			}
			sectionLength += 1
			continue
		default:
			if confirmedLength != 0 {
				confirmed = append(confirmed, []int{confirmedLength, cI})
				cI = -1
				confirmedLength = 0
			}
			if sectionLength != 0 {
				sections = append(sections, []int{sectionLength, sI})
				sectionLength = 0
				sI = -1
			}
		}
	}
	if confirmedLength != 0 {
		confirmed = append(confirmed, []int{confirmedLength, cI})
		cI = -1
		confirmedLength = 0
	}
	if sectionLength != 0 {
		sections = append(sections, []int{sectionLength, sI})
	}
	return sections, confirmed
}

func defineSpringRanges(springs []int, overall int) []*BrokenSprings {
	ranges := make([]*BrokenSprings, len(springs))
	min := 0
	total := len(springs) - 1
	for _, s := range springs {
		total += s
	}
	upper := total
	for i, s := range springs {
		max := overall - (upper)
		if max < 0 {
			max = overall - 1
		}
		ranges[i] = &BrokenSprings{
			count:  s,
			min:    min,
			max:    max,
			index:  i,
			starts: []int{},
		}

		min += s + 1
		upper -= s + 1
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})
	ranges[len(ranges)-1].max = overall - ranges[len(ranges)-1].count
	return ranges
}

func defineSpringRangesBetter(s string, springs []int) []*BrokenSprings {
	ranges := make([]*BrokenSprings, len(springs))
	totalSpringSize := len(springs) - 1
	for _, size := range springs {
		totalSpringSize += size
	}
	min := 0
	overall := len(s)
	for i, size := range springs {
		spr := &BrokenSprings{
			count: size,
		}
		for true {
			c := s[min]
			if c == '.' {
				min += 1
				continue
			}
			invalid := false
			for i := 1; i < spr.count; i++ {
				c = s[min+i]
				if c == '.' {
					min += 1
					invalid = true
					break
				}
			}
			if invalid {
				continue
			}
			break
		}
		spr.min = min
		spr.max = overall - (totalSpringSize)
		for spr.max != spr.min {
			c := s[spr.max]
			if c == '.' {
				spr.max -= 1
				continue
			}
			invalid := false
			for i := 1; i < spr.count; i++ {
				c = s[spr.max+i]
				if c == '.' {
					spr.max -= 1
					invalid = true
					break
				}
			}
			if invalid {
				continue
			}
			break
		}
		spr.dof = spr.max - spr.min
		ranges[i] = spr
		min = spr.min + spr.count + 1
		totalSpringSize -= (spr.count + 1)
	}
	return ranges
}

func recursive(source string, springs []*BrokenSprings) string {
	required := make([][]int, len(source))
	requiredSprings := make([]bool, len(springs))

	for i, c := range source {
		if c == '#' {
			for j, s := range springs {
				if s.min <= i && i <= s.max+s.count-1 {
					required[i] = append(required[i], j)
				}
				if s.min > i {
					break
				}
			}
			r := required[i]
			if len(r) == 1 {
				requiredSprings[r[0]] = true
				newMin := i - springs[r[0]].count + 1
				if newMin > springs[r[0]].min {
					springs[r[0]].min = newMin
				}
				if i < springs[r[0]].max {
					springs[r[0]].max = i
				}
			}
		}
	}

	for i, yes := range requiredSprings {
		if !yes {
			continue
		}
		for j := springs[i].min; j <= springs[i].max; j++ {
			valid, complete := true, true
			for k := 0; k < springs[i].count; k++ {
				if source[j+k] == '.' {
					valid = false
				}
				if source[j+k] == '?' {
					complete = false
				}
			}
			if valid && complete {
				springs[i].starts = append(springs[i].starts, j)
				break
			}
			if valid {
				springs[i].starts = append(springs[i].starts, j)
			}
		}
	}

	outputArray := make([]string, len(source))
	output := source
	for _, c := range springs {
		if len(c.starts) == 1 {
			for i, s := range output {
				if i == c.starts[0]-1 {
					outputArray[i] = "."
					continue
				}
				if i == c.starts[0]+c.count {
					outputArray[i] = "."
					continue
				}
				if i >= c.starts[0] && i < c.starts[0]+c.count {
					outputArray[i] = "#"
					continue
				}
				outputArray[i] = string(s)
			}
			output = strings.Join(outputArray, "")
		}
	}
	return strings.Trim(output, ".")
}

func buildPossibilityArrayV2(source string, springs []*BrokenSprings) int {
	for i, c := range source {
		if c == '.' {
			for _, s := range springs {
				if i < s.min || i > s.max {
					continue
				}
				if s.min == i {
					s.min = i + 1
				}
			}
		}
	}
	for i := len(source) - 1; i >= 0; i-- {
		c := source[i]
		if c == '.' {
			for _, s := range springs {
				if i < s.min || i > s.max {
					continue
				}
				if s.max == i {
					s.max = i - 1
				}
			}
		}
	}

	for i, s := range springs[:len(springs)-1] {
		for _, r := range springs[i+1:] {
			if r.min <= s.min+s.count+1 {
				r.min = s.min + s.count + 1
			}
			if s.max >= r.max-2 {
				s.max = r.max - 2
			}
		}
	}

	required := make([][]int, len(source))
	for i, c := range source {
		if c == '#' {
			for j, s := range springs {
				if s.min <= i && i <= s.max+s.count-1 {
					required[i] = append(required[i], j)
				}
				if s.min > i {
					break
				}
			}
		}
	}

	requiredSprings := make([]bool, len(springs))
	for i, r := range required {
		if len(r) == 1 {
			requiredSprings[r[0]] = true
			newMin := i - springs[r[0]].count + 1
			if newMin > springs[r[0]].min {
				springs[r[0]].min = newMin
			}
		}
		if len(r) > 1 {
			for _, spr := range r {
				j := i
				start := 2
				invalid := false
				for _, springa := range springs[spr:] {
					for j < i+start+springa.count {
						if j >= len(required) {
							invalid = true
							break
						}
						if required[j] != nil && !utils.Contains(springa.index, required[j]) {
							invalid = true
							break
						}
						j += 1
					}
					start += springa.count
				}
				if !invalid {
					requiredSprings[spr] = true
					newMin := i - springs[spr].count + 1
					if newMin > springs[spr].min {
						springs[spr].min = newMin
					}
				}
			}
		}
	}

	for i, s := range springs[:len(springs)-1] {
		for _, r := range springs[i+1:] {
			if r.min <= s.min+s.count+1 {
				r.min = s.min + s.count + 1
			}
			if s.max >= r.max-2 {
				s.max = r.max - 2
			}
		}
	}

	for i, yes := range requiredSprings {
		if !yes {
			continue
		}
		for j := springs[i].min; j <= springs[i].max; j++ {
			valid, complete := true, true
			for k := 0; k < springs[i].count; k++ {
				if source[j+k] == '.' {
					valid = false
				}
				if source[j+k] == '?' {
					complete = false
				}
			}
			if valid && complete {
				springs[i].starts = []int{j}
				break
			}
			if valid {
				springs[i].starts = append(springs[i].starts, j)
			}
		}
	}

	possibilities := make([][]int, len(source))
	for i, c := range source {
		switch c {
		case '#', '?':
			possibilities[i] = []int{}
			for j, s := range springs {
				if s.min > i {
					break
				}
				if j > 0 {
					if s.min <= springs[j-1].min+springs[j-1].count+1 {
						s.min = springs[j-1].min + springs[j-1].count + 1
					}
				}
				if s.isPossibility(i) {
					/*if i > 0 && utils.Contains(j, required[i-1]) {
						continue
					}
					if i < len(source)-1 && utils.Contains(j, required[i+1]) {
						continue
					}*/

					possibilities[i] = append(possibilities[i], j)
				}
			}
		}
	}

	counter := len(springs) - 1
	for i := len(source) - 1; i >= 0; i-- {
		if possibilities[i] == nil {
			continue
		}
		if utils.Contains(counter, possibilities[i]) {
			if i <= springs[counter].max {
				springs[counter].max = i
				counter -= 1
			}
			continue
		}

	}

	for i, s := range springs[:len(springs)-1] {
		for _, r := range springs[i+1:] {
			if r.min <= s.min+s.count+1 {
				r.min = s.min + s.count + 1
			}
			if s.max >= r.max-2 {
				s.max = r.max - 2
			}
		}
	}

	for i, c := range source {
		switch c {
		case '#', '?':
			possibilities[i] = []int{}
			for j, s := range springs {
				if s.min > i {
					break
				}
				if j > 0 {
					if s.min <= springs[j-1].min+springs[j-1].count+1 {
						s.min = springs[j-1].min + springs[j-1].count + 1
					}
				}
				if s.isPossibility(i) {

					possibilities[i] = append(possibilities[i], j)
				}
			}
		}
	}

	for j, s := range springs {
		if len(s.starts) > 0 {
			continue
		}
		if s.min == s.max {
			s.starts = []int{s.min}
			continue
		}
		for i, p := range possibilities {
			if utils.Contains(j, p) {
				valid := true
				for k := 0; k < s.count; k++ {
					if source[i+k] == '.' {
						valid = false
					}
				}
				if valid {
					s.starts = append(s.starts, i)

				}
			}
		}
	}

	for _, s := range springs {
		s.print()
	}
	val := evaluate(springs, source)
	fmt.Println(val)
	return val
}

func evaluate(springs []*BrokenSprings, source string) int {
	result := make([]bool, len(source))
	complex := []*BrokenSprings{}
	for _, s := range springs {
		if len(s.starts) != 1 {
			complex = append(complex, s)
			continue
		}
		for i := 0; i < s.count; i++ {
			result[i+s.starts[0]] = true
		}
	}
	if len(complex) == 0 {
		return 1
	}
	return 0 //evaluateComplex(complex, result, source)
}

/*

func evaluateComplex(springs []*BrokenSprings, result []bool, source string) int {
	count := 0
	candidate := springs[0]
	if len(springs) == 1 {
		for _, start := range candidate.starts {
			isValid, _ := isValidCondition(start, candidate.count, result, source)
			if !isValid {
				continue
			}
			count += 1
		}
		return count
	}

	for _, start := range candidate.starts {
		isValid, resultCopy := isValidCondition(start, candidate.count, result, source)
		if !isValid {
			continue
		}
		count += evaluateComplex(springs[1:], resultCopy, source)
	}
	return count
}
*/

func buildPossibilityArray(source string, springs []*BrokenSprings) {
	required := make([][]int, len(source))
	for i, c := range source {
		if c == '.' {
			for _, s := range springs {
				if i < s.min {
					break
				}
				if s.min == i {
					s.min = i + 1
				}
				if s.max == i {
					s.max = i - 1
				}
			}
		}
	}

	for _, s := range springs {
		s.print()
	}

	for i, c := range source {
		if c == '#' {
			for j, s := range springs {
				if s.min <= i && i <= s.max+s.count {
					required[i] = append(required[i], j)
				}
				if s.min > i {
					break
				}
			}
		}
	}
	for i, r := range required {
		if len(r) == 1 {
			newMin := i - springs[r[0]].count + 1
			if newMin > springs[r[0]].min {
				springs[r[0]].min = newMin
			}
		}
	}

	possibilities := computePossibilities(springs, source)

	counter := len(springs) - 1
	for i := len(source) - 1; i >= 0; i-- {
		if possibilities[i] == nil {
			continue
		}
		if utils.Contains(counter, possibilities[i]) {
			if i < springs[counter].max {
				springs[counter].max = i
				counter -= 1
			}
			continue
		}

	}

	possibilities = computePossibilities(springs, source)

	/*
		counter := 0
		for i, p := range possibilities {
			if p != nil {
				if len(p) > 1 && utils.Contains(counter, p) {
					for j, s := range springs[counter+1:] {
						if
					}
				}
			}
		}*/
	/*j := 0
	for i, s := range springs {
		if possibilities[j] != nil {
			if possibilities[j]
		}
	}*/

}

func computePossibilities(springs []*BrokenSprings, source string) [][]int {
	for i, s := range springs[:len(springs)-1] {
		for _, r := range springs[i+1:] {
			if r.min <= s.min+s.count+1 {
				r.min = s.min + s.count + 1
			}
			if s.max >= r.max-2 {
				s.max = r.max - 2
			}
		}
	}

	possibilities := make([][]int, len(source))
	for i, c := range source {
		switch c {
		case '#', '?':
			possibilities[i] = []int{}
			for j, s := range springs {
				if s.min > i {
					break
				}
				if j > 0 {
					if s.min <= springs[j-1].min+springs[j-1].count+1 {
						s.min = springs[j-1].min + springs[j-1].count + 1
					}
				}
				if s.min <= i && i <= s.max {
					/*if i > 0 && utils.Contains(j, required[i-1]) {
						continue
					}
					if i < len(source)-1 && utils.Contains(j, required[i+1]) {
						continue
					}*/

					possibilities[i] = append(possibilities[i], j)
				}
			}
		}
	}
	return possibilities
}

func evaluateConfirmed(springs []*BrokenSprings, confirmed [][]int) {
	sorted := make([]*BrokenSprings, len(springs))
	for i, s := range springs {
		sorted[i] = s
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].count > sorted[j].count
	})

}

func solvex(springs []*BrokenSprings, sections [][]int, source string) (bool, int) {
	sorted := make([]*BrokenSprings, len(springs))
	for i, s := range springs {
		sorted[i] = s
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].count > sorted[j].count
	})
	totalOptions := 0
	if len(springs) == 1 {

	}
	for i, s := range sections {
		if s[0] >= sorted[0].count {
			// Section is big enough
			for j := 0; j+sorted[0].count <= s[0]; j++ {
				leftSections := sections[0:i]
				if j > 1 {
					leftSections = append(leftSections, []int{j, s[1]})
				}
				if len(leftSections) > 0 && len(springs[:sorted[0].index]) > 0 {
					isValid, options := solvex(springs[:sorted[0].index], leftSections, source[:s[1]+j])
					if !isValid {
						continue
					}
					totalOptions += options
				}
				rightSections := sections[i+1:]
				if s[0]-(j+sorted[0].count) > 1 {
					rightSections = append([][]int{{s[0] - (j + sorted[0].count), s[1] + (j + sorted[0].count)}}, rightSections...)
				}
				if len(rightSections) > 0 && len(springs[sorted[0].index+1:]) > 0 {
					isValid, options := solvex(springs[sorted[0].index+1:], rightSections, source[s[1]+(j+sorted[0].count):])
					if !isValid {
						continue
					}
					totalOptions += options
				}
			}
		}
	}
	return false, 0
}

func solve(s string, nums []int, database map[string]int) int {
	key := getKey(s, nums)
	fmt.Println(key, s, nums, "solve")
	if len(nums) == 0 || len(s) == 0 {
		return 1
	}
	val, ok := database[key]
	if ok {
		return val
	}
	springs := defineSpringRangesBetter(s, nums)
	factor := 1
	hasDefinite := false
	for i, spring := range springs {
		if spring.dof == 0 {
			hasDefinite = true
			if i != 0 {
				factor *= solve(s[:spring.min-1], nums[:i], database)
			}
			if i != len(springs)-1 {
				factor *= solve(s[spring.max+spring.count+1:], nums[i+1:], database)
			}
			break
		}
	}
	if hasDefinite {
		database[key] = factor
		return factor
	}
	chars := recursive(s, springs)
	if chars != s {
		return solve(chars, nums, database)
	}
	val = bruteforce(s, springs, []int{}, nums)
	database[key] = val
	return val
}

func bruteforce(source string, springs []*BrokenSprings, choices, expected []int) int {
	if len(choices) == len(springs) {
		val := check(source, springs, choices, expected)
		return val
	}
	candidate := springs[len(choices)]
	total := 0
	minimum := candidate.min
	if len(choices) != 0 {
		minimum = choices[len(choices)-1] + 2
	}
	for i := minimum; i <= candidate.max; i++ {
		choicesCopy := append(choices, i)
		total += bruteforce(source, springs, choicesCopy, expected)
	}
	return total
}

func check(source string, springs []*BrokenSprings, choices, expected []int) int {
	result := make([]bool, len(source))
	for i, s := range springs {
		startIndex := choices[i]
		endIndex := startIndex + s.count
		for j := startIndex; j < endIndex; j++ {
			if source[j] == '.' {
				return 0
			}
			if result[j] {
				return 0
			}
			result[j] = true
		}
		if startIndex > 0 {
			if result[startIndex-1] {
				return 0
			}
		}
		if endIndex < len(source)-1 {
			if result[endIndex+1] {
				return 0
			}
		}
	}
	counter := 0
	value := 0
	for i, s := range source {
		if s == '#' && !result[i] {
			return 0
		}
		if result[i] {
			value += 1
		} else {
			if value != 0 {
				if expected[counter] != value {
					return 0
				}
				counter += 1
				value = 0
			}
		}
	}
	return 1
}

func getKey(s string, nums []int) string {
	return fmt.Sprintf("%s:%v", s, nums)
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a := ArrangeSprings(lines)
	return strconv.Itoa(a), "B"
}
