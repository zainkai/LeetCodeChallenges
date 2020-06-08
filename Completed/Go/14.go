func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    } else if len(strs) == 1 {
        return strs[0]
    }
    
    idx := 0
    topLoop:
    for {
        for i := 1; i < len(strs); i++ {
            word1, word2 := strs[i-1], strs[i]
            if idx >= len(word1) || idx >= len(word2) {
                break topLoop
            }
            
            if word1[idx] != word2[idx] {
                break topLoop
            }
        }
        idx++
    }
    
    return strs[0][0:idx]
}
