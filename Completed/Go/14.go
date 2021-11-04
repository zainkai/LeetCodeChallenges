func longestCommonPrefix(strs []string) string {
    i := 0
    
    topLoop:
    for ;i < len(strs[0]); i++ {
        ch := strs[0][i]
        for _, str := range strs {
            if i >= len(str) || str[i] != ch { 
                break topLoop
            }
        }
    }
    
    return strs[0][:i]
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}
