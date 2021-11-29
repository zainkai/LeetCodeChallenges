func singleNumber(nums []int) int {
    res := 0
    for _, elm := range nums {
        res = res ^ elm
    }
    
    return res
}
