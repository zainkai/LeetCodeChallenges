func rob(nums []int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		a := 0
		b := 0

		if i-2 >= 0 {
			a = nums[i-2]
		}
		if i-3 >= 0 {
			b = nums[i-3]
		}

		nums[i] += max(a, b)

		result = max(result, nums[i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}