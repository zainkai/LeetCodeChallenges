import (
    "strings"
    "sort"
)

func frequencySort(s string) string {
    arr := []string{}
    freq := map[string]int{}
    // get freqs O(N)
    for _, r := range s {
        str := string(r)
        freq[str]++
        
        arr = append(arr, str)
    }
    
    // heapify O(N log N)
    sort.Slice(arr, func(i, j int) bool {
        a, b := arr[i], arr[j]
        if freq[a] == freq[b] {
            return a > b
        }
        return freq[a] > freq[b]
    })
    
    // O(N)
    return strings.Join(arr, "")
}
