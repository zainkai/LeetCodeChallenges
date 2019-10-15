import "strings"

func reverseWords(s string) string {
	tokens := strings.Split(s, " ")

	for i, token := range tokens {
		tokens[i] = reverse(token)
	}

	return strings.Join(tokens, " ")
}

func reverse(s string) string {
	l := 0
	r := len(s) - 1
	result := []rune(s)

	for l < r {
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}

	return string(result)
}