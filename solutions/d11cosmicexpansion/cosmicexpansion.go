package d11cosmicexpansion

import (
	"advent/solutions/utils"
	"math"
	"sort"
	"strconv"
)

func ExpandGalaxy(lines []string, increment int) int {
	galaxies := buildGalaxy(lines, increment)
	return galaxyPaths(galaxies)
}

// buildGalaxy takes the input galaxy and progressively increments coordinates
func buildGalaxy(lines []string, increment int) [][]int {
	columns := make([]bool, len(lines[0]))
	for i := range columns {
		columns[i] = true
	}
	galaxies := [][]int{}

	// first iterate across x axis, looking for vacant rows
	// increment all y coordinates by the currentYIncrement
	// if there is a vacant row, increment the total increment to represent the added rows
	//
	// Also flag if a column ever contains a galaxy
	currentYIncrement := 0
	for i, l := range lines {
		vacant := true
		for j, c := range l {
			if c == '#' {
				vacant = false
				columns[j] = false
				galaxies = append(galaxies, []int{j, i + currentYIncrement})
			}
		}
		if vacant {
			currentYIncrement += increment
		}
	}

	// sort the galaxies by x coordinate
	// then simultaneously iterate over the original columns and galaxies sorted by original x coord
	// increment all x coords by the currentXIncrement
	// if the column was vacant, increment the total increment
	sort.Slice(galaxies, func(i, j int) bool {
		return galaxies[i][0] < galaxies[j][0]
	})
	currentXIncrement := 0
	j := 0
	for i, c := range columns {
		nextGalaxy := galaxies[j]
		for nextGalaxy[0] <= i {
			nextGalaxy[0] += currentXIncrement
			j += 1
			if j >= len(galaxies) {
				break
			}
			nextGalaxy = galaxies[j]
		}
		if c {
			currentXIncrement += increment
			continue
		}
	}
	return galaxies
}

// galaxyPaths iterates across all galaxy pairs and totals distance
func galaxyPaths(galaxies [][]int) int {
	total := 0
	for i, g := range galaxies {
		for j := i; j < len(galaxies); j++ {
			h := galaxies[j]
			total += manhattan(g, h)
		}
	}
	return total
}

// manhattan computes the manhattan distance of two points
func manhattan(a, b []int) int {
	return int(math.Abs(float64(b[0]-a[0])) + math.Abs(float64(b[1]-a[1])))
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a, b := ExpandGalaxy(lines, 1), ExpandGalaxy(lines, 999999)
	return strconv.Itoa(a), strconv.Itoa(b)
}
