func twoSum(numbers []int, target int) []int {
	m := map[int]int{}

	for i, v := range numbers {
		if j, ok := m[v]; ok {
			return []int{j + 1, i + 1}
		}

		comp := target - v
		m[comp] = i
	}

	return []int{-1, -1}
}