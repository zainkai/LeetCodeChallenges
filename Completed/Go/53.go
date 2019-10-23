func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		t := nums[i] + nums[i-1]
		nums[i] = max(nums[i], t)

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