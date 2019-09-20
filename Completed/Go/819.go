import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := strArrToMap(banned)
	paragraph = clean(paragraph)
	tokens := map[string]int{}

	result := ""
	resultCount := 0
	for _, str := range strings.Split(paragraph, " ") {
		if bannedMap[str] > 0 || str == "" {
			continue
		}

		tokens[str] += 1

		if tokens[str] > resultCount {
			result = str
			resultCount = tokens[str]
		}
	}

	return result
}

func clean(s string) string {
	return strings.Map(func(c rune) rune {
		if unicode.IsPunct(c) {
			return ' '
		}
		return unicode.ToLower(c)
	}, s)
}

func strArrToMap(arr []string) map[string]int {
	res := map[string]int{}

	for _, str := range arr {
		res[str] += 1
	}

	return res
}