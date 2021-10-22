func areOccurrencesEqual(s string) bool {
    if len(s) == 0 {
        return true
    }
    
    counts := map[byte]int{}
    for _, r := range s {
        counts[byte(r)]++
    }
    
    tmp := s[0]
    for k := range counts {
        if counts[k] != counts[tmp] {
            return false
        } 
    }
    
    return true
}
