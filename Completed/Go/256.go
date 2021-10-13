type frame struct {
    color, house int
}

func minCost(costs [][]int) int {
    f := frame{-1, 0}
    memo := map[frame]int{}
    
    return helper(f, memo, costs)
}

func helper(f frame, memo map[frame]int, costs [][]int) int {
    if res, ok := memo[f]; ok {
        return res
    } else if f.house == len(costs) {
        return 0
    }
    
    memo[f] = 1<< 63-1
    for color, cost := range costs[f.house] {
        if color == f.color {
            continue
        }
        memo[f] = min(
            memo[f],
            helper(frame{color, f.house+1}, memo, costs)+cost,
        )
    }
    
    return memo[f]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
