package LeetCode

func firstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, char := range s {
		m[char] += 1
	}

	for i, char := range s {
		if m[char] == 1 {
			return i
		}
	}

	return -1

}
