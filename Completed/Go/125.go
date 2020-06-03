import (
    "strings"
)

func isPalindrome(s string) bool {
    i, j := 0, len(s)-1
    for i < j {
        if !isChar(s[i]) {
            i++
        } else if !isChar(s[j]) {
            j--
        } else if isEq(s[i], s[j]) {
            i++
            j--
        } else {
            return false
        }
    }
    
    return true
}

func isChar(b byte) bool {
    if b >= 'a' && b <= 'z' {
        return true
    } else if b >= 'A' && b <= 'Z' {
        return true
    } else if b >= '0' && b <= '9' {
        return true
    }
    return false
}

func isEq(a, b byte) bool {
    i := string(a)
    j := string(b)
    
    return strings.EqualFold(i,j)
}
