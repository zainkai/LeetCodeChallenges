func buildArray(nums []int) []int {
    res := make([]int, len(nums))
    for i, v := range nums {
        res[i] = nums[v]
    }
    
    return res
}
