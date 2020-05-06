func minCostClimbingStairs(costs []int) int {
	memo := map[int]int{}

	return min(
		helper(costs, 0, memo),
		helper(costs, 1, memo),
	)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func helper(costs []int, pos int, memo map[int]int) int {
	if cost, ok := memo[pos]; ok {
		return cost
	} else if pos >= len(costs) {
		return 0
    }

	cost := costs[pos]
	memo[pos] = cost + min(
		helper(costs, pos+1, memo),
		helper(costs, pos+2, memo))

	return memo[pos]
}


