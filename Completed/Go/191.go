func hammingWeight(num uint32) int {
    mask := uint32(1)
    res := 0
    
    for mask > 0 {
        if mask & num >= 1 {
            res++
        }
        mask <<= 1
    }
    return res
}
