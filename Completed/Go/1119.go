import (
	"strings"
)

var vowelMap = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func isVowelRune(r rune) bool {
	return vowelMap[r]
}

func removeVowels(S string) string {
	result := strings.Builder{}

	for _, r := range S {
		if isVowelRune(r) {
			continue
		}

		result.WriteRune(r)
	}

	return result.String()
}