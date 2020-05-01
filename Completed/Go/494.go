type frame struct {
    pos, sum int
}

func findTargetSumWays(nums []int, S int) int {
    memo := map[frame]int{}
    
    return helper(memo, nums, frame{0,0}, S)
}

func helper(memo map[frame]int, nums []int, f frame, S int) int {
    if v, ok := memo[f]; ok {
        return v
    } else if f.pos == len(nums) {
        if f.sum == S {
            return 1
        }
        return 0
    } else {
        memo[f] += helper(memo, nums, frame{f.pos+1, f.sum+nums[f.pos]}, S)
        memo[f] += helper(memo, nums, frame{f.pos+1, f.sum-nums[f.pos]}, S)
    }
    return memo[f]
}