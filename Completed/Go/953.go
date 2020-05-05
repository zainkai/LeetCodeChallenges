func isAlienSorted(words []string, order string) bool {
    m := getOrderMap(order)
    
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]
        if res := cmp(word1, word2, m); res == 1 {
            return false
        }
    }
    return true
}

// 0 - word1 == word2
// -1 - word1 < word2
// 1 - word2 < word1
func cmp(word1, word2 string, m map[byte]int) int {
    maxLen := len(word1)
    if len(word2) > maxLen {
        maxLen = len(word2)
    }
    for i := 0; i < maxLen; i++ {
        a := byte(0)
        if i < len(word1) {
            a = word1[i]
        }
        b := byte(0)
        if i < len(word2) {
            b = word2[i]
        }
        
        if m[a] == m[b] {
            continue
        } else if m[a] < m[b] {
            return -1
        } else {
            return 1
        }
    }
    return 0
}

func getOrderMap(order string) map[byte]int {
    m := map[byte]int{}
    for i, r := range order {
        m[byte(r)] = i
    }
    return m
}