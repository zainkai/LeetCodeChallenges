func tribonacci(n int) int {
	m := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}

	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2] + m[i-3]
	}

	return m[n]
}