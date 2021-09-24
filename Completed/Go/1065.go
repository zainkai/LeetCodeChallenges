import "sort"

func indexPairs(text string, words []string) [][]int {
	
	// O(words log words)
    sort.Slice(words, func(i,j int) bool {
        return len(words[i]) < len(words[j])
    })
    
    res := [][]int{}
	// O(words * text)
    for i := range text {
        for _, word := range words {
            j := i + len(word)
            if j > len(text) {
                continue
            }
            tmp := text[i:j]
            if tmp == word {
                res = append(res, []int{i,j-1})
            }
        }
    }
    return res
}
