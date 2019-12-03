func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	} else if len(nums) == 1 {
		return true
	}

	j := 0
	for i := 0; i <= j; i++ {
		if j >= len(nums)-1 {
			return true
		}

		j = max(j, i+nums[i])
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}