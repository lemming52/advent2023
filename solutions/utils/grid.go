package utils

func NESW2DNeighbours(x, y, xMax, yMax int) [][]int {
	coords := [][]int{
		{x, y - 1},
		{x, y + 1},
		{x - 1, y},
		{x + 1, y},
	}
	correct := [][]int{}
	for _, xy := range coords {
		if xy[0] >= 0 && xy[1] >= 0 && xy[0] <= xMax && xy[1] <= yMax {
			correct = append(correct, xy)
		}
	}
	return correct
}

func All2DNeighbours(x, y, xMax, yMax int) [][]int {
	coords := [][]int{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		{x - 1, y},
		{x + 1, y},
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
	correct := [][]int{}
	for _, xy := range coords {
		if xy[0] >= 0 && xy[1] >= 0 && xy[0] <= xMax && xy[1] <= yMax {
			correct = append(correct, xy)
		}
	}
	return correct
}
