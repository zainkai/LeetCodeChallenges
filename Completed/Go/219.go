
func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]bool{}
	for i, v := range nums {
		if m[v] {
			return true
		}

		m[v] = true

		if i >= k {
			key := nums[i-k]
			m[key] = false
		}
	}

	return false
}