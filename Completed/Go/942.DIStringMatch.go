func diStringMatch(S string) []int {
	var low, high int = 0, len(S)
	result := []int{}
	
	for _, val := range S {
			if val == 'I' {
					result = append(result, low)
					low++
			} else {
					result = append(result, high)
					high--
			}
	}
	
	result = append(result, low)
	
	return result
}