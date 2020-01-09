const ( // tri state bool
	unvisited    = 0
	visitedTrue  = 1
	visitedFalse = 2
)

func wordBreak(s string, wordDict []string) bool {
	dict := createDict(wordDict)
	mem := make([]int, len(s))

	return helper(0, s, dict, mem)
}

func helper(pos int, s string, dict map[byte][]string, mem []int) bool {
	if pos >= len(s) {
		return true
	} else if mem[pos] == visitedTrue {
		return true
	} else if mem[pos] == visitedFalse {
		return false
	}

	for _, word := range dict[s[pos]] {

		// does each char in word fit in s
		newPos := pos + len(word)
		if len(s) < newPos {
			continue
		}

		t := string(s[pos:newPos])
		if t == word && helper(pos+len(word), s, dict, mem) {
			mem[pos] = visitedTrue
			return true
		} else {
			mem[pos] = visitedFalse
		}
	}

	return false
}

func createDict(wordDict []string) map[byte][]string {
	dict := make(map[byte][]string)

	for _, word := range wordDict {
		char := word[0]
		if _, ok := dict[char]; !ok {
			dict[char] = []string{word}
		} else {
			dict[char] = append(dict[char], word)
		}
	}

	return dict
}