import "sort"
// O(N * klogk)
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}

	for _, str := range strs {

		key := string(sortBytes(str))
		if _, ok := m[key]; !ok {
			m[key] = []string{str}
		} else {
			m[key] = append(m[key], str)
		}
	}

	result := [][]string{}
	for _, sArr := range m {
		result = append(result, sArr)
	}

	return result
}

func sortBytes(s string) string {
	b := []byte(s)

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	return string(b)
}

// O(N *K)
// func groupAnagrams(words []string) [][]string {
// 	m := map[[26]int][]string{}

// 	for _, word := range words {
// 		arr := wordToCharCount(word)
// 		m[arr] = append(m[arr], word)
// 	}

// 	res := [][]string{}
// 	for _, val := range m {
// 		res = append(res, val)
// 	}

// 	return res
// }

// func wordToCharCount(word string) [26]int {
// 	res := [26]int{}

// 	for _, r := range word {
// 		res[r-'a']++
// 	}
// 	return res
// }