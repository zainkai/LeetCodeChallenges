func productExceptSelf(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    
    res := make([]int, len(nums))
    
    L := make([]int, len(nums))
    tmp := 1
    for i := 0; i < len(nums); i++ {
        tmp *= nums[i]
        L[i] = tmp
    }
    
    R:= make([]int, len(nums))
    tmp = 1
    for i := len(nums)-1; i >=0; i-- {
        tmp *= nums[i]
        R[i] = tmp
    }
    
    for i := range nums {
        left := 1
        if i-1 >= 0 {
            left = L[i-1]
        }
        
        right := 1
        if i+1 < len(nums) {
            right = R[i+1]
        }
        
        res[i] = left * right
    }
    return res
}
