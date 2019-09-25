func hammingDistance(x int, y int) int {
	count := 0
	for x != 0 || y != 0 {
		if x%2 != y%2 {
			count++
		}
		x >>= 1
		y >>= 1
	}
	return count
}