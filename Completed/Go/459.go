func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return true
	} else if len(s) == 1 {
		return false
	}

	for i := 1; i <= len(s)/2; i++ {
		a := string(s[:i])

		for j := i; j <= len(s)-i; j += i {
			b := string(s[j : j+i])

			if a != b {
				break
			} else if j == len(s)-i {
				return true
			}
		}
	}
	return false
}
