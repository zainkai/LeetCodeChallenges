func canConstruct(ransomNote, magazine string) bool {
	runeCounts := make(map[rune]int)
	for _, r := range magazine {
		runeCounts[r]++
	}
	
	for _, r := range ransomNote {
		if count, ok := runeCount[r]; !ok {
			return false
		} else if count == 0 {
			return false
		} else {
			runeCount[r]--
		}
	}
	return true
}
	
	