func isPowerOfTwo(n int) bool {
    if n == 1 {
        return true
    }
    
    mask := 2
    for mask < n {
        mask <<= 1
    }
    return n == mask
}