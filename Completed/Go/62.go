type point struct {
    x,y int
}

func uniquePaths(m int, n int) int {
    memo := map[point]int{}
    start := point{0,0}
    
    return helper(memo, start, m-1, n-1)
}

func helper(memo map[point]int, curr point, m, n int) int {
    if curr.x > m || curr.y > n {
        return 0
    } else if curr.x == m && curr.y == n {
        return 1
    } else if val, ok := memo[curr]; ok {
        return val
    }
    
    right := point{curr.x+1, curr.y}
    memo[curr] += helper(memo, right, m, n)
    
    down := point{curr.x, curr.y+1}
    memo[curr] += helper(memo, down, m, n)
    
    return memo[curr]
}