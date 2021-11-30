func largest1BorderedSquare(grid [][]int) int {
    res := 0
    
    for i := 0; i < len(grid);i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 0 {
                continue
            }
            
            x, y := i, j
            tmp:= 1
            for isInbounds(grid, x, y) {
                if grid[x][y] == 1 && checkBoundary(grid, i, j, x, y) {
                    res = max(res, tmp*tmp)
                }
                x++
                y++
                tmp++
            }
            
            
        }
    }
    
    return res
}

func checkBoundary(grid [][]int, i, j, x, y int) bool {    
    for a := i; a < x; a++ {
        if grid[a][j] == 0 || grid[a][y] == 0 {
            return false
        }
    }
    
    for a := j; a < y; a++ {
        if grid[i][a] == 0 || grid[x][a] == 0 {
            return false
        }
    }
    return true
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func isInbounds(grid [][]int, a, b int) bool {
    if a < 0 || b < 0 {
        return false
    } else if a >= len(grid) || b >= len(grid[a]) {
        return false
    }
    return true
}

