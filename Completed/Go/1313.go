func decompressRLElist(nums []int) []int {
    res := make([]int, 0)
    
    for i := 1; i < len(nums); i += 2 {
        freq := nums[i-1]
        val := nums[i]
        
        for j := 0; j < freq; j++ {
            res = append(res, val)
        }
    }
    
    return res
}