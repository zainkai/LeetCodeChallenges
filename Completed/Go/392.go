func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	} else if len(t) == 0 {
		return false
	}

	j := 0
	for i := 0; i < len(t); i++ {
		sByte := s[j]
		tByte := t[i]
		if sByte != tByte {
			continue
		}

		j++
		if j >= len(s) {
			return true
		}
	}

	return false
}