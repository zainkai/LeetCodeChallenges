func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		if nums[0] == k {
			return 1
		} else {
			return 0
		}
	}

	sum, s := make([]int, len(nums)), 0
	for i, val := range nums {
		s += val
		sum[i] = s
	}

	res := 0
	m := map[int]int{
		0: 1,
	}
	for _, val := range sum {
		if v, ok := m[val-k]; ok {
			res += v
		}

		m[val] += 1
	}

	return res
}