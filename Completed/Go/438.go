func findAnagrams(s string, p string) []int {
    pMap := getCharArr(p)
    res := []int{}
    
    for i := 0; i < len(s) - len(p)+1; i++ {
        tmp := s[i:i+len(p)]
        
        if pMap == getCharArr(tmp) {
            res = append(res, i)
        }
    }
    
    return res
}

func getCharArr(s string) [26]int {
    res := [26]int{}
    for _, elm := range s {
        idx := int(elm - 'a')
        res[idx]++
    }
    
    return res
}
