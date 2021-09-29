func sortArrayByParityII(nums []int) []int {
    
    for i := 0; i < len(nums); i++ {
        j := i+1
        if i % 2 == 0 && nums[i] % 2 != 0  {
            for nums[j] % 2 != 0 {
                j++
            }
            nums[i], nums[j] = nums[j], nums[i]
        } else if i % 2 != 0 && nums[i] % 2 == 0 {
            for nums[j] % 2 == 0 {
                j++
            }
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    
    
    return nums
}
