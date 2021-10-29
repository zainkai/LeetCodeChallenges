func titleToNumber(columnTitle string) int {
    digit := 1
    res := 0
    
    for i := len(columnTitle)-1; i >= 0; i-- {
        charToInt := int(columnTitle[i] - 'A')+1
        
        res += (digit * charToInt)
        digit *= 26
    }
    
    return res
}
