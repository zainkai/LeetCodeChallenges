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