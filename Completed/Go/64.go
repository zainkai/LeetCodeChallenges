func minPathSum(grid [][]int) int {
	n, m := len(grid), 0
	if n > 0 {
		m = len(grid[0])
	}

	for i, row := range grid {
		for j, val := range row {
			if i == 0 && j == 0 {
				continue
			}

			left := 1<<63 - 1
			if isValidSpace(n, m, i, j-1) {
				left = grid[i][j-1] + val
			}

			top := 1<<63 - 1
			if isValidSpace(n, m, i-1, j) {
				top = grid[i-1][j] + val
			}

			grid[i][j] = min(left, top)
		}
	}

	if n > 0 && m > 0 {
		return grid[n-1][m-1]
	}
	return 0
}

func isValidSpace(n, m, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	} else if i >= n || j >= m {
		return false
	} else {
		return true
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}