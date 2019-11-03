import "sort"

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