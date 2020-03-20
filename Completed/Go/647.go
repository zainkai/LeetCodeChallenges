// O(N^2), O(1)
func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += countPalindromes(s, i, i)   // count for odd length palindromes
		count += countPalindromes(s, i, i+1) // count for even length palindromes
	}

	return count
}

func countPalindromes(s string, L, R int) int {
	i, j := L, R

	count := 0
	for i >= 0 && j < len(s) {
		if s[i] == s[j] {
			count++
		} else {
			break
		}
		i--
		j++
	}

	return count
}