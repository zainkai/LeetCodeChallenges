func shortestCompletingWord(licensePlate string, words []string) string {
	licenseMap := StrToArr(licensePlate)
	res := "                                                   "

	for _, word := range words {
		wordMap := StrToArr(word)
		if compareWordMappings(licenseMap, wordMap) && len(word) < len(res) {
			res = word
		}
	}

	return res
}

func StrToArr(s string) []int {
	m := make([]int, 26)
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			lowerR := r - 'A'
			m[lowerR]++
		} else if r >= 'a' && r <= 'z' {
			m[r-'a']++
		}
	}

	return m
}

func compareWordMappings(licence, word []int) bool {
	for key, val := range licence {
		if word[key] < val {
			return false
		}
	}
	return true
}