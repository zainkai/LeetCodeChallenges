func reverseBits(num uint32) uint32 {
    mask := uint32(1)
    rMask := uint32(1 << 31)
    res := uint32(0)
    
    for mask > 0 {
        if mask & num > 0 {
            res = res | rMask
        }
        mask <<= 1
        rMask >>= 1
    }
    
    return res
}
