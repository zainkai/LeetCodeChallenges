func maxIncreaseKeepingSkyline(grid [][]int) int {
	rowMax := make([]int, len(grid))
	colMax := make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			rowMax[i] = max(grid[i][j], rowMax[i])
			colMax[j] = max(grid[i][j], colMax[j])
		}
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			m := min(rowMax[i], colMax[j])
			g := grid[i][j]

			result += m - g
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}