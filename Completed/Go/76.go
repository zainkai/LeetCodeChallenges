func minWindow(s string, t string) string {    
    tMap := getCharCount(t)
    sMap := map[byte]int{}
    
    i, j := 0, -1
    l, r := 0, 1<<63-1
    for j <= len(s) {
        if GTECounts(sMap, tMap) {
            if r-l > j-i {
                l = i
                r = j
            }
            
            b := s[i]
            sMap[b]--
            i++
        } else {
            j++
            if j < len(s) {
                b := s[j]
                sMap[b]++
            }
        }
    }
    
    if r < 0 || r > len(s) {
        return ""
    }
    return s[l:r+1]
}

func getCharCount(t string) map[byte]int {
    counts := map[byte]int{}
    for i := 0; i < len(t); i++ {
        b := t[i]
        counts[b]++
    }
    return counts
}

// O(26)
func GTECounts(s, t map[byte]int) bool {
    for b, val := range t {
        if val > s[b] {
            return false
        }
    }
    return true
}