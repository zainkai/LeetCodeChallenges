func findLengthOfLCIS(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    
    prev := -1 << 63
    
    res := 1
    temp := 0
    for _, val := range nums {
        if val <= prev {
            prev = -1 << 63
            temp = 0
        }
        
        prev = val
        temp += 1

        res = max(temp, res)
    }
    
    return res
}

func max (a, b int) int {
    if a > b {
        return a
    }
    return b
}