func missingNumber(nums []int) int {
	res := len(nums)
	for i, v := range nums {
		res = res ^ v ^ i
	}

	return res
}