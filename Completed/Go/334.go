func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	min1, min2 := 1<<63-1, 1<<63-1
	for _, val := range nums {
		if val < min1 {
			min1 = val
		} else if val > min1 && val < min2 {
			min2 = val
		} else if val > min1 && val > min2 {
			return true
		}
	}

	return false
}