func maxProduct(nums []int) int {
    curr := 1
    maxNum := nums[0]
    for _, val := range nums {
        if curr == 0 {
            curr = 1
        }
        
        curr *= val
        maxNum = max(maxNum, curr)
    }
    
    curr = 1
    for i := len(nums)-1; i >= 0; i-- {
        if curr == 0 {
            curr = 1
        }
        curr *= nums[i]
        maxNum = max(maxNum, curr)
    }
    
    return maxNum
}

func max(nums ...int) int {
    res := nums[0]
    for _, num := range nums {
        if num > res {
            res = num
        }
    }
    return res
}