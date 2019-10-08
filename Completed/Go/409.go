func longestPalindrome(s string) int {
	charCount := map[rune]int{}
	result := 0
	for _, r := range s {
		charCount[r] += 1
		if charCount[r]%2 == 0 {
			result += 2
		}
	}

	if result != len(s) {
		result += 1
	}
	return result
}