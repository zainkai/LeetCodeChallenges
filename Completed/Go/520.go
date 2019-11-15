func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}

	for i := 1; i < len(word)-1; i++ {
		a := isCap(word[i])
		b := isCap(word[i+1])
		if (a && !b) || (!a && b) {
			return false
		}
	}

	c := isCap(word[0])
	d := isCap(word[1])
	return (c && d) || (!c && !d) || (c && !d)
}

func isCap(b byte) bool {
	return b <= 'Z' && b >= 'A'
}