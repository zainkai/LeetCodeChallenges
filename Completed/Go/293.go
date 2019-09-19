func generatePossibleNextMoves(s string) []string {
	results := []string{}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '-' || s[i+1] == '-' {
			continue
		}

		newStr := s[:i] + "--" + s[i+2:]
		results = append(results, newStr)
	}

	return results
}