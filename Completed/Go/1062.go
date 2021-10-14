//naive

func longestRepeatingSubstring(s string) int {
    freq := map[string]int{}
    for i := 0; i < len(s); i++ {
        for j := i; j < len(s); j++ {
            sub := s[i:j+1]
            freq[sub]++
        }
    }
    
    maxSubLen := 0
    for key, val := range freq {
        if len(key) > maxSubLen && val >= 2 {
            maxSubLen = len(key)
        }
    }
    return maxSubLen
}
