func largestUniqueNumber(A []int) int {
	m := map[int]int{}

	for _, n := range A {
		m[n] += 1
	}

	result := -1
	for key, count := range m {
		if count != 1 {
			continue
		}

		if result < key {
			result = key
		}
	}

	return result
}