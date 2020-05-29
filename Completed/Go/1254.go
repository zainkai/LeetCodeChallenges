const (
    water = 1
    land  = 0
)

func closedIsland(grid [][]int) int {
    for i := 0; i < len(grid[0]); i++ {
        fillLand(grid, 0, i)
        fillLand(grid, len(grid)-1, i)
    }
    for i := 0; i < len(grid); i++ {
        fillLand(grid, i, 0)
        fillLand(grid, i, len(grid[0])-1)
    }
    
    
    islandCount := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == land {
                islandCount++
                fillLand(grid, i, j)
            }
            
        } 
    }
    
    return islandCount
}

func fillLand(grid [][]int, x, y int) {
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
        return
    } else if grid[x][y] == water {
        return
    } 
    
    grid[x][y] = water
    moves := [][]int{{-1,0}, {1,0}, {0,1}, {0,-1}}
    for _, mov := range moves  {
        fillLand(grid, x+mov[0], y +mov[1])
    }
}