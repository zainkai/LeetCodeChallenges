func fib(N int) int {
	m := map[int]int{
		0: 0,
		1: 1,
	}

	for i := 2; i <= N; i++ {
		m[i] = m[i-1] + m[i-2]
	}

	return m[N]
}