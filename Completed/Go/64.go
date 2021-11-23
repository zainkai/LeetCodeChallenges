type coord struct {
    x,y int
}

func minPathSum(grid [][]int) int {
    memo := map[coord]int{
        coord{ len(grid)-1, len(grid[0])-1 }:grid[len(grid)-1][len(grid[0])-1],
        
    }
    return helper(grid, coord{0,0}, memo)
}

func helper(grid [][]int, curr coord, memo map[coord]int) int {
    if res, ok := memo[curr]; ok {
        return res
    } else if curr.x >= len(grid) || curr.y >= len(grid[curr.x]) {
        return 1<<63-1
    }
    
    memo[curr] = 1<<63-1
    if res := helper(grid, coord{curr.x+1, curr.y}, memo); res != 1<<63-1 {
        memo[curr] = min(memo[curr], res)
    }
    if res := helper(grid, coord{curr.x, curr.y+1}, memo); res != 1<<63-1 {
        memo[curr] = min(memo[curr], res)
    }
    
    memo[curr] += grid[curr.x][curr.y]
    return memo[curr]
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}
