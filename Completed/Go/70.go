func climbStairs(n int) int {
	m := map[int]int{
		0: 1,
		1: 1,
	}

	for i := 2; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}

	return m[n]
}

func climbStairs(n int) int {
    memo := map[int]int{
        0: 1,
        1: 1,
    }
    
    return helper(memo, n)
}

func helper(memo map[int]int, n int) int {
    if paths, ok := memo[n]; ok {
        return paths
    }
    
    memo[n] = helper(memo, n-1) + helper(memo, n-2)
    return memo[n]
}
