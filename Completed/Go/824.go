import "strings"

var vowels = map[byte]bool{
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

func toGoatLatin(S string) string {

	result := []string{}
	tokens := strings.Split(S, " ")
	for i, token := range tokens {
		str := ""
		if isFirstVowel(token) {
			str = token + maGen(i+2)
		} else {
			f := token[0]
			str = token[1:] + string(f) + maGen(i+2)
		}

		result = append(result, str)
	}

	return strings.Join(result, " ")
}

func isFirstVowel(s string) bool {
	if len(s) == 0 {
		return false
	}

	f := s[0]
	return vowels[f]
}

func maGen(i int) string {
	return "m" + strings.Repeat("a", i)
}