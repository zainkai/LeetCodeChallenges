func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	for _, v := range nums {
		key := abs(v) - 1
		if nums[key] < 0 {
			continue
		}

		nums[key] = -nums[key]
	}

	for i, v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}