func getMaximumGold(grid [][]int) int {
	maxGold := 0
	for i := range grid {
		for j := range grid[i] {
			maxGold = max(maxGold, dfs(grid, i, j))
		}
	}

	return maxGold
}

var directions [][]int = [][]int{[]int{0, 1}, []int{0, -1}, []int{-1, 0}, []int{1, 0}}

func dfs(grid [][]int, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
		return 0
	} else if grid[x][y] == 0 {
		return 0
	}

	gold, maxTrav := grid[x][y], 0
	grid[x][y] = 0

	for _, p := range directions {
		maxTrav = max(maxTrav, dfs(grid, x+p[0], y+p[1]))
	}
	grid[x][y] = gold

	return gold + maxTrav
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}