func sortColors(nums []int)  {
    c1 := 0
    c2 := 0
    
    for _, c := range nums {
        if c == 0 {
            c1++
        } else if c == 1 {
            c2++
        }
    }
    
    for i := range nums {
        if c1 > 0 {
            nums[i] = 0
            c1--
        } else if c2 > 0 {
            nums[i] = 1
            c2--
        } else {
            nums[i] = 2
        }
    }
}