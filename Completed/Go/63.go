const (
	EMPTY    int = 0
	OBSTACLE int = 1
)

func uniquePathsWithObstacles(grid [][]int) int {
	if grid[0][0] == OBSTACLE || grid[len(grid)-1][len(grid[0])-1] == OBSTACLE {
		return 0
	}

	grid[0][0] = 2
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == OBSTACLE {
				continue
			}

			if isValid(grid, i+1, j) {
				grid[i+1][j] += grid[i][j]
			}
			if isValid(grid, i, j+1) {
				grid[i][j+1] += grid[i][j]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1] / 2

}

func isValid(grid [][]int, i int, j int) bool {
	if i < 0 || j < 0 {
		return false
	} else if i > len(grid)-1 || j > len(grid[0])-1 {
		return false
	} else if grid[i][j] == OBSTACLE {
		return false
	}

	return true
}

