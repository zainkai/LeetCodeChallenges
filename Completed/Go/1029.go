type frame struct {
    idx, bucket int
}

func twoCitySchedCost(costs [][]int) int {
    memo := map[frame]int{}
    return helper(costs, frame{0,0}, memo)
}

func helper(costs [][]int, f frame, memo map[frame]int) int {
    if res, ok := memo[f]; ok {
        return res
    } else if f.idx == len(costs) && f.bucket == len(costs)/2{
        return 0
    } else if f.idx >= len(costs) {
        return 1<<62
    }
    
    memo[f] = min(
        helper(costs, frame{f.idx+1, f.bucket+1}, memo) +costs[f.idx][0],
        helper(costs, frame{f.idx+1, f.bucket}, memo) +costs[f.idx][1],
    )
    return memo[f]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
