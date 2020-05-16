func reverseWords(s []byte)  {
    reverseBytes(s, 0, len(s)-1)
    
    i, j := 0,0
    for i < len(s) && j < len(s) {
        if s[j] != ' ' {
            j++
        } else {
            reverseBytes(s, i, j-1)
            i = j+1
            j = i
        }
    }
    reverseBytes(s, i, j-1)
}

func reverseBytes(s []byte, start, end int) {
    if end <= start {
        return
    }
    
    for start < end {
        s[start], s[end] = s[end], s[start]
        
        start++
        end--
    }
    
}