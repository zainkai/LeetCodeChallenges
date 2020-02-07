func longestStrChain(words []string) int {
	memo := make(map[string]int)
	maxDepth := 0

	wordMap := map[string]bool{}
	for _, word := range words {
		wordMap[word] = true
	}

	for _, word := range words {
		maxDepth = max(maxDepth, helper(word, memo, wordMap))
	}

	return maxDepth
}

func helper(word string, memo map[string]int, wordMap map[string]bool) int {
	if len(word) == 1 {
		return 1
	} else if _, ok := memo[word]; ok {
		return memo[word]
	}

	maxSoFar := 1
	for _, pred := range removeAtIndexString(word, wordMap) {
		maxSoFar = max(maxSoFar, helper(pred, memo, wordMap)+1)
	}

	memo[word] = maxSoFar
	return maxSoFar
}

func removeAtIndexString(in string, wordMap map[string]bool) []string {
	strs := []string{}
	for i := range in {
		if str := in[:i] + in[i+1:]; wordMap[str] {
			strs = append(strs, str)
		}
	}

	return strs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}