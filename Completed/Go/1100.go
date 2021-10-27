func numKLenSubstrNoRepeats(s string, k int) int {
    if k > len(s) {
        return 0
    }
    cMap := genCharMap(s[:k])
    
    l, r := 0, k-1
    res := 0
    for r < len(s) {
        if isUniqueChars(cMap) {
            res++
        }
        if r+1 >= len(s) {
            break
        }
        
        cMap[s[l]]--
        l++
        
        r++
        cMap[s[r]]++
    }
    
    return res
}

func genCharMap(s string) map[byte]int {
    res := map[byte]int{}
    for i := 0; i < len(s); i++ {
        b := byte(s[i])
        res[b]++
    }
    return res
}

func isUniqueChars(m map[byte]int) bool {
    for _, v := range m {
        if v > 1 {
            return false
        }
    }
    return true
}
