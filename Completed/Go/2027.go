func minimumMoves(s string) int {
    res := 0
    for i := 0; i< len(s);i++ {
        if s[i] == 'X' {
            res++
            i+= 2
        }
    }
    return res
}
