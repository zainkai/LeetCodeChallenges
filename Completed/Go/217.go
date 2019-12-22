func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}

	for _, key := range nums {
		if _, ok := m[key]; ok {
			return true
		}
		m[key] = struct{}{}
	}

	return false
}