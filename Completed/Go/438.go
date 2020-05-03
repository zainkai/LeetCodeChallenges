func findAnagrams(s string, p string) []int {
    pMap, pCount := map[byte]int{}, 0
    for _, r := range p {
        pMap[byte(r)]++
        pCount++
    }
    
    result := []int{}
    sMap, sCount := map[byte]int{}, 0
    for i := 0; i < len(s); i++ {
        if sCount <= pCount {
            sMap[s[i]]++
            sCount++
        }
        if sCount > pCount {
            char := s[i-len(p)]
            sMap[char]--
            if sMap[char] == 0 {
                delete(sMap, char)
            }
            sCount--
        }
        
        if compareMaps(pMap, sMap) {
            result = append(result, i-len(p)+1)
        }
    }
    return result
}

func compareMaps(a, b map[byte]int) bool {
    if len(a) != len(b) {
        return false
    }
    
    for key := range a {
        if b[key] != a[key] {
            return false
        }
    }
    return true
}