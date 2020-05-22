import (
    "sort"
    "strings"
)

func frequencySort(s string) string {
    charCount := map[rune]int{}
    
    uniques := []rune{}
    for _, r := range s {
        if _, ok := charCount[r]; !ok {
            uniques = append(uniques, r)
        }
        charCount[r]++
    }
    sort.Slice(uniques, func(i,j int) bool {
        a, b := uniques[i], uniques[j]
        return charCount[a] > charCount[b]
    })
    
    sb := strings.Builder{}
    for _, r := range uniques {
        for i := 0; i < charCount[r]; i++ {
            sb.WriteRune(r)
        }
    }
    
    return sb.String()
}
