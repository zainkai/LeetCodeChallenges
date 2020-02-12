import (
	"unicode"
)

func licenseKeyFormatting(S string, K int) string {
	result, k := make([]byte, 0), 0

	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		} else {
			result = append(result, toUpper(S[i]))
		}

		k++
		if k == K {
			k = 0
			result = append(result, '-')
		}
	}

	res := string(reverseBytes(result))
	if len(res) > 0 && res[0] == '-' {
		return res[1:]
	}
	return res
}

func toUpper(b byte) byte {
	return byte(unicode.ToUpper(rune(b)))
}

func reverseBytes(arr []byte) []byte {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return arr
}