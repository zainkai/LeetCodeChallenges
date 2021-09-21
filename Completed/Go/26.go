func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    
    anchor := 0
    runner := 0
    
    for runner < len(nums) {        
        a := nums[anchor]
        b := nums[runner]
        
        if a != b{
            anchor++
            nums[anchor] = nums[runner]
        }
        runner++
    }
    
    return anchor+1
}
