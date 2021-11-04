type WordDistance struct {
    indexes map[string][]int
}


func Constructor(wordsDict []string) WordDistance {
    indexes := map[string][]int{}
    for i, word := range wordsDict {
        indexes[word] = append(indexes[word], i)
    }
    
    return WordDistance{
        indexes,
    }
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    res := 1<<63-1
    
    for _, i := range this.indexes[word1] {
        for _, j := range this.indexes[word2] {
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


/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */
