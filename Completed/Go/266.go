func canPermutePalindrome(s string) bool {
    counts := map[rune]int{}
    odds := 0
    for _, r := range s {
        counts[r]++
        
        if counts[r] % 2 == 0 {
            odds--
        } else {
            odds++
        }
    }
    
    return odds <= 1
}
