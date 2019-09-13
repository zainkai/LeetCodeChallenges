const (
	one  = "1"
	zero = "0"
)

func addBinary(a string, b string) string {
	aArr := reverse(a)
	bArr := reverse(b)

	strLen := max(len(a), len(b))

	result := ""
	carry := false
	for i := 0; i < strLen; i++ {
		var x, y byte = '0', '0'
		if i < len(a) {
			x = a[i]
		}
		if i < len(b) {
			y = b[i]
		}

		if x == '1' && y == '1' {
			if carry {
				result = "1" + result
			} else {
				carry = true
				result = "0" + result
			}
		} else if x == '0' && y == '0' {
			if carry {
				carry = false
				result = "1" + result
			} else {
				result = "0" + result
			}
		} else {
			if carry {
				result = "0" + result
			} else {
				result = "1" + result
			}
		}
	}

	if carry {
		result = "1" + result
	}

	return result
}

func reverse(s string) []byte {
	chars := []byte(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return chars
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
