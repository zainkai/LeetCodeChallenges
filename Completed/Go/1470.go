func shuffle(nums []int, n int) []int {
    a := 0
    b := len(nums)/2
    
    res := []int{}
    
    for i := 0; i < len(nums)/2; i ++ {
        res = append(res, nums[a], nums[b])
        a++
        b++
    }
    
    return res
}
