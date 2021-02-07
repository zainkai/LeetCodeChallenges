func shortestToChar(s string, c byte) []int {
    result := make([]int, len(s))
    tmp := 1<<62-1
    
    if s[0] == c {
        tmp = 0
    }
    result[0] = tmp
    
    for i := 1; i < len(s); i++ {
        ch := s[i]
        if ch == c {
            tmp = 0
        } else {
            tmp = result[i-1] +1
        }
        result[i] = tmp
    }
    
    tmp = 1<< 62-1
    if s[len(s)-1] == c {
        tmp = 0
    }
    result[len(s)-1] = min(result[len(s)-1], tmp)
    for i := len(s)-2; i >= 0; i-- {
        ch := s[i]
        if ch == c {
            tmp = 0
        } else {
            tmp = result[i+1]+1
        }
        result[i] = min(result[i], tmp)
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
