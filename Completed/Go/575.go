func distributeCandies(candies []int) int {
	m := map[int]int{}

	for _, candy := range candies {
		m[candy] += 1
	}

	a := len(m)
	b := len(candies) / 2
	if a < b {
		return a
	}
	return b
}