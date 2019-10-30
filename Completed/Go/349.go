func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, v := range nums1 {
		m[v] += 1
	}

	result := []int{}
	for _, v := range nums2 {
		if m[v] > 0 {
			result = append(result, v)
			delete(m, v)
		}
	}

	return result
}