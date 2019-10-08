func canPermutePalindrome(s string) bool {
	m := map[rune]int{}
	result := 0
	for _, r := range s {
		m[r]++
		if m[r]%2 == 0 {
			result += 2
		}
	}

	return result >= len(s)-1
}