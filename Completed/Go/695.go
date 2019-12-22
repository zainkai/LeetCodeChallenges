const (
	WATER          int = 0
	UNVISITED_LAND int = 1
	VISITED_LAND   int = -1
)

func maxAreaOfIsland(grid [][]int) int {
	maxAreaIsland := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == UNVISITED_LAND {
				areaOfIsland := getIslandArea(grid, i, j)

				if areaOfIsland > maxAreaIsland {
					maxAreaIsland = areaOfIsland
				}
			}
		}
	}
	return maxAreaIsland
}

func getIslandArea(grid [][]int, row, col int) int {
	if row < 0 || col < 0 {
		return 0
	} else if row >= len(grid) || col >= len(grid[row]) {
		return 0
	} else if grid[row][col] == WATER || grid[row][col] == VISITED_LAND {
		return 0
	}

	grid[row][col] = VISITED_LAND

	return (1 + getIslandArea(grid, row+1, col) + getIslandArea(grid, row-1, col) + getIslandArea(grid, row, col+1) + getIslandArea(grid, row, col-1))
}