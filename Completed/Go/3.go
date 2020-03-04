func lengthOfLongestSubstring(s string) int {
    l, r, maxLen := 0, 0, 0
    m := map[byte]int{}   
    
    for r < len(s) {
        b := s[r]
        if m[b] == 0 {
            m[b]++
            r++
            maxLen = max(maxLen, r- l)
            continue
        }
        // else
        for l < len(s) {
            a := s[l]
            m[a]--
            l++
            if a == b {
                break
            }
        }
    }
    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}