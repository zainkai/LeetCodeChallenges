func orangesRotting(grid [][]int) int {
	ans := -1

	countFresh := 0
	wasChanged := true
	for wasChanged {
		wasChanged, countFresh = step(grid)
		ans++
	}

	if countFresh == 0 {
		return ans
	}
	return -1
}

const (
	empty         = 0
	fresh         = 1
	rotten        = 2
	freshToRotten = 3
)

func step(grid [][]int) (bool, int) {
	countFresh := 0
	wasChanged := false

	// detect Rotten routine
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == rotten {
				changeAdj(grid, i, j)
			}
		}
	}
	// change freshToRotten back to Rotten
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == freshToRotten {
				grid[i][j] = rotten
				wasChanged = true
			} else if grid[i][j] == fresh {
				countFresh++
			}
		}
	}

	return wasChanged, countFresh
}

func changeAdj(grid [][]int, i, j int) {
	mods := [][]int{
		{-1, 0}, {1, 0}, {0, 1}, {0, -1},
	}
	for _, mod := range mods {
		x := i + mod[0]
		y := j + mod[1]
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
			continue
		} else if grid[x][y] == fresh {
			grid[x][y] = freshToRotten
		}
	}
}