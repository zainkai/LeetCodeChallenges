

func numMatchingSubseq(S string, words []string) int {
    wordMap := map[string]int{}
    out := 0
    for _, w := range words {
        wordMap[w]++
    }
    
    for word, val := range wordMap {
        if isSubSequence(S, word) {
            out += val
        }
    }
    return out
}

func isSubSequence(s, t string) bool {
    j := 0
    for i := 0; i < len(s) && j < len(t); i++ {
        if s[i] == t[j] {
            j++
        }
    }
    return j == len(t)
}
