type Point struct {
    x, y int
}

func maximalSquare(grid [][]byte) int {
    ans := 0
    memo := map[Point]int{}
    
    for x, row := range grid {
        for y := range row {
            if grid[x][y] == '1' {
               ans = max(ans, dfs(memo, grid, x, y))   
            } 
        }
    }
    return ans * ans
}

func max (a,b int) int {
    if a > b {
        return a
    }
    return b
}

func min (vals... int) int {
    res := vals[0]
    for _, val := range vals {
        if val < res {
            res = val
        }
    }
    return res
}

func dfs(memo map[Point]int, grid [][]byte, x ,y int) int {
    p := Point{x,y}
    if val, ok := memo[p]; ok {
        return val
    } else if x >= len(grid) || y >= len(grid[x]) {
        return 0
    } else if grid[x][y] == '0' {
        return 0
    }
    
    down := dfs(memo, grid, x+1, y)
    right := dfs(memo, grid, x, y+1)
    if down >= 1 && right >= 1 {
        memo[p] = min(down, right, dfs(memo, grid, x+1, y+1)) +1
    } else {
        memo[p] = 1
    }
    return memo[p]
}