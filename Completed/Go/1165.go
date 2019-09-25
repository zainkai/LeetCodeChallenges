func calculateTime(keyboard string, word string) int {
	m := map[rune]int{}
	for i, key := range keyboard {
		m[key] = i
	}

	lastIdx := 0
	result := 0
	for _, c := range word {
		result += abs(m[c] - lastIdx)
		lastIdx = m[c]
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}