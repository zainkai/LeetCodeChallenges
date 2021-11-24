type point struct {
    x,y int
}

func minFallingPathSum(matrix [][]int) int {
    if len(matrix) == 0 {
        return 0
    }
    
    res := 1<<63-1
    memo := map[point]int{}
    for y := range matrix[0] {
        res = min(res, helper(
            matrix,
            point{0,y},
            memo,
        ))
    }
    return res
}

func helper(matrix [][]int, p point, memo map[point]int) int {
    if p.x == len(matrix) {
        return 0
    } else if isOutOfBounds(matrix, p) {
        return 1<<63-1
    } else if res, ok := memo[p]; ok {
        return res
    }
    
    memo[p] = 1<<63-1
    memo[p] = min(memo[p], helper(matrix, point{p.x+1, p.y}, memo))
    memo[p] = min(memo[p], helper(matrix, point{p.x+1, p.y-1}, memo))
    memo[p] = min(memo[p], helper(matrix, point{p.x+1, p.y+1}, memo))
    
    memo[p] += matrix[p.x][p.y]
    return memo[p]
} 

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func isOutOfBounds(g [][]int, p point) bool {
    return p.x <0 || p.y < 0 || p.x >= len(g) || p.y >= len(g[p.x])
}
