func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	x, y := 0, 0

	for i := 0; i < len(s)-1; i++ {
		if a, b := checkPalin(s, i, i); (b - a) >= (y - x) { // odd length palin
			x, y = a, b
		}
		if a, b := checkPalin(s, i, i+1); (b - a) >= (y - x) { // even length palin
			x, y = a, b
		}
	}

	return string(s[x : y+1])
}

func checkPalin(s string, l, r int) (int, int) {
	x, y := 0, 0

	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			break
		}

		x, y = l, r
		l--
		r++
	}

	return x, y
}