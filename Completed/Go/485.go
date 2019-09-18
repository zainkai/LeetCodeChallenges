func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	count := 0
	for _, n := range nums {
		if n == 0 {
			count = 0
		} else {
			count += 1
		}

		if count > result {
			result = count
		}
	}

	return result
}