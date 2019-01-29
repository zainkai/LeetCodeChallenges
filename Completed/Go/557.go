import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
			words[i] = reverseString(word)
	}
	
	return strings.Join(words, " ")
}

func reverseString(s string) string {
	r := []rune(s)
for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
	r[i], r[j] = r[j], r[i]
}
return string(r)
}