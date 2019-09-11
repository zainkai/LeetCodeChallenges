func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, num := range nums {
		val, ok := m[num]
		if ok && i != val {
			return []int{i, val}
		}

		comp := target - num
		m[comp] = i
	}

	return []int{-1, -1}
}