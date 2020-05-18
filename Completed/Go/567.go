func checkInclusion(p string, s string) bool {
    charMap := getMap(p)
    
    l, r := 0,0
    
    temp := [26]int{}
    for r < len(s) {
        if r-l < len(p) {
            idx := s[r] - 'a'
            temp[idx]++
            r++
        } else {
            lIdx := s[l] -'a'
            temp[lIdx]--
            l++
            
            rIdx := s[r]-'a'
            temp[rIdx]++
            r++
        }
        
        
        if temp == charMap {
            return true
        }
    }
    
    return false
}


func getMap(p string) [26]int {
    charMap := [26]int{}
    
    for _, r := range p {
        idx := r - 'a'
        charMap[idx]++
    }
    
    return charMap
}