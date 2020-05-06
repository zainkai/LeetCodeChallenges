func combinationSum4(nums []int, target int) int {
    memo := map[int]int{
        0: 1,
    }
    
    return dfs(nums, target, memo)
}

func dfs(nums []int, target int, memo map[int]int) int {
    if target == 0 {
        return 1
    } else if target < 0 {
        return 0
    } else if val, ok := memo[target]; ok  {
        return val
    }
    
    for _, val := range nums {
        memo[target] += dfs(nums, target-val, memo)
    }
    
    return memo[target]
}
