func minCostII(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	memo := initMemo(costs)
	minCosts := 1<<63 - 1

	for color := range costs[0] {
		minCosts = min(minCosts, dfs(0, color, costs, memo))
	}

	return minCosts
}

func dfs(house int, prevColor int, costs [][]int, memo [][]int) int {
	if house == len(costs) {
		return 0
	}

	if memo[house][prevColor] != -1 {
		return memo[house][prevColor]
	}

	minPath := 1<<63 - 1
	for color := 0; color < len(costs[house]); color++ {
		if color == prevColor {
			continue
		}

		minPath = min(minPath, dfs(house+1, color, costs, memo))
	}

	if minPath == 1<<63-1 {
		minPath = 0
	}
	memo[house][prevColor] = costs[house][prevColor] + minPath

	return memo[house][prevColor]
}

//------------------------- utils
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func initMemo(costs [][]int) [][]int {
	memo := make([][]int, len(costs))
	for i := 0; i < len(costs); i++ {
		for j := 0; j < len(costs[i]); j++ {
			memo[i] = append(memo[i], -1)
		}
	}

	return memo
}