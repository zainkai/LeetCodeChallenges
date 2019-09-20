func convertToTitle(n int) string {
	result := ""
	for n > 0 {
		char, i := NumToColumn(n)

		result = char + result
		n = i
	}

	return result
}

func NumToColumn(n int) (string, int) {
	n -= 1

	i := int(n % 26)
	remainder := int(n / 26)

	return string('A' + i), remainder
}