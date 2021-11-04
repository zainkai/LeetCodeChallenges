func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    idx1 := []int{}
    idx2 := []int{}
    
    for i, word := range wordsDict {
        if word == word1 {
            idx1 = append(idx1, i)
        } else if word == word2 {
            idx2 = append(idx2, i)
        }
    }
    
    res := 1<<63-1
    for _, i := range idx1 {
        for _, j := range idx2 {
            val := abs(i-j)
            res = min(res, val)
        }
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
