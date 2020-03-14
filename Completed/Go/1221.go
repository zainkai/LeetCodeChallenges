func balancedStringSplit(s string) (pairs int) {
	lcounter := 0

	for _, r := range s {
		if r == 'L' {
			lcounter++
		} else {
			lcounter--
		}

		if lcounter == 0 {
			pairs++
		}
	}

	return pairs
}