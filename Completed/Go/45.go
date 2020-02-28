// dynamic programming TLE
func jump(nums []int) int {
	memo := make([]int, len(nums))
	dfs(nums, memo, 0, 0)

	return memo[len(memo)-1]
}

func dfs(nums []int, memo []int, pos int, jumps int) {
	if pos >= len(nums) {
		return
	} else if memo[pos] != 0 && memo[pos] < jumps {
		return
	}

	memo[pos] = jumps
	for i := pos + 1; i <= pos+nums[pos]; i++ {
		dfs(nums, memo, i, jumps+1)
	}
}

// greedy
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	jumps := 1

	currJumpLen := nums[0]
	end := nums[0]
	for i := 1; i < len(nums); i++ {
		if currJumpLen < i {
			currJumpLen = end
			jumps++
		}

		if i+nums[i] > end {
			end = i + nums[i]
		}
	}

	return jumps
}
