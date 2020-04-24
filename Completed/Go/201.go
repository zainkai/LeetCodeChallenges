func rangeBitwiseAnd(m int, n int) int {
    shift := 0
    for m != n {
        m >>=1
        n >>= 1
        
        shift++
    }
    return m << shift
}

// brute force
// func rangeBitwiseAnd(m int, n int) int {
//     res := m
//     for i := m+1; i <= n; i++ {
//         res = res & i
//     }
//     return res
// }
