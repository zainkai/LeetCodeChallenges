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

// recursive
func rob(nums []int) int {
    res := 0
    memo := map[int]int{}
    
    for i := range nums {
        res = max(
            res,
            helper(memo, nums, i),
        )
    }
    return res
}

func helper(memo map[int]int, nums []int, idx int) int {
    if idx >= len(nums) {
        memo[idx] = 0
        return 0
    } else if res, ok := memo[idx]; ok {
        return res
    }
    
    memo[idx] = nums[idx] + max(
        helper(memo, nums, idx+2),
        helper(memo, nums, idx+3),
    )
    
    return memo[idx]
}
