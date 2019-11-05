func findTheDifference(s string, t string) byte {
	result := rune(0)

	for _, r := range s {
		result ^= r
	}
	for _, r := range t {
		result ^= r
	}

	return byte(result)
}