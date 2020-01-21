func productExceptSelf(nums []int) []int {
	if len(nums) < 0 {
		return nums
	}

	output, product := make([]int, len(nums)), 1
	for i := len(nums) - 1; i >= 0; i-- {
		product *= nums[i]
		output[i] = product
	}

	product = 1
	for i, val := range nums {
		rVal := 1
		if i+1 < len(nums) {
			rVal = output[i+1]
		}

		output[i] = product * rVal
		product *= val
	}

	return output
}