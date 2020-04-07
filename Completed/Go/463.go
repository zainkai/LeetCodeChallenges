func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				perimeter += checkPerimeter(grid, i, j)
			}
		}
	}
	return perimeter
}

func checkPerimeter(grid [][]int, i, j int) int {
	perimeter := 0

	mod := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, coor := range mod {
		x, y := coor[0], coor[1]
		i += x
		j += y

		if i < 0 || i >= len(grid) {
			perimeter++
		} else if j < 0 || j >= len(grid[i]) {
			perimeter++
		} else if grid[i][j] == 0 {
			perimeter++
		}

		i -= x
		j -= y
	}
	return perimeter
}