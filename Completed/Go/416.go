type frame struct {
    i, target int
}

func canPartition(nums []int) bool {
    memo := map[frame]bool{}
    
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if sum % 2 == 1 {
        return false
    }
    return helper(nums, frame{0, (sum/2)}, memo)
}

func helper(nums []int, f frame, memo map[frame]bool) bool {
    if res, ok := memo[f]; ok {
        return res
    } else if f.target == 0 {
        memo[f] = true
        return memo[f]
    } else if f.i == len(nums) || f.target < 0 {
        memo[f] = false
        return memo[f]
    }
    
    memo[f] = memo[f] || helper(nums, frame{f.i+1,f.target-nums[f.i]}, memo)
    memo[f] = memo[f] || helper(nums, frame{f.i+1, f.target}, memo)
    
    return memo[f]
}


/// Brute force DP
/*
type frame struct {
    i, s1, s2 int
}

func canPartition(nums []int) bool {
    memo := map[frame]bool{}
    
    return helper(nums, frame{0, 0, 0}, memo)
}

func helper(nums []int, f frame, memo map[frame]bool) bool {
    if res, ok := memo[f]; ok {
        return res
    } else if f.i == len(nums) {
        memo[f] = f.s1 == f.s2
        return memo[f]
    }
    
    val := nums[f.i]
    memo[f] = memo[f] || helper(nums, frame{f.i+1, f.s1+val, f.s2}, memo)
    memo[f] = memo[f] || helper(nums, frame{f.i+1, f.s1, f.s2+val}, memo)
    
    return memo[f]
}

*/
