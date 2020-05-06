// SLOW

func lengthOfLIS(nums []int) int {
    memo := map[int]int{}
    
    res := helper(nums, -1, memo)
    return res
}

func helper(n []int, pos int, memo map[int]int) int {
    if val, ok := memo[pos]; ok {
        return val
    } 
    
    num := -1 << 63
    if pos != -1 {
        num = n[pos]
    }
    
    memo[pos]= 0
    for i := pos+1; i < len(n); i++ {
        if n[i] > num {
            memo[pos] = max(helper(n, i, memo)+1, memo[pos])
        }
    }
    
    return memo[pos]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}