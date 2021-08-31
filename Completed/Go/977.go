func sortedSquares(nums []int) []int {
    result := make([]int, len(nums))
    resCount := len(nums)-1
    
    lo, hi := 0, len(nums)-1
    for lo <= hi {
        lo2 := nums[lo] * nums[lo]
        hi2 := nums[hi] * nums[hi]
        
        if lo2 > hi2 {
            result[resCount] = lo2
            lo++
        } else {
            result[resCount] = hi2
            hi--
        }
        resCount--
    }
    return result
}
