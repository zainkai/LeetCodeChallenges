const (
	land  byte = '1'
	water byte = '0'
)

func numIslands(grid [][]byte) int {
	numIslands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == land {
				numIslands++
				convertLandToWater(grid, i, j)
			}
		}
	}

	return numIslands
}

func convertLandToWater(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return
	}

	if grid[i][j] == water {
		return
	}

	grid[i][j] = water

	convertLandToWater(grid, i+1, j)
	convertLandToWater(grid, i-1, j)
	convertLandToWater(grid, i, j+1)
	convertLandToWater(grid, i, j-1)
}

