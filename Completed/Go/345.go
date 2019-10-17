import "strings"

func reverseVowels(s string) string {
	vowels := strings.Builder{}
	for _, r := range s {
		if isVowelRune(r) {
			vowels.WriteRune(r)
		}
	}

	result := strings.Builder{}
	vowelStr := vowels.String()
	i := len(vowelStr) - 1

	for _, r := range s {
		if isVowelRune(r) {
			b := vowelStr[i]
			result.WriteByte(b)
			i--
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}

var vowelMap = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func isVowelRune(r rune) bool {
	return vowelMap[r]
}