func numJewelsInStones(J string, S string) int {
	jewelTypes := map[rune]bool{}
	for _, jewel := range J {
		jewelTypes[jewel] = true
	}

	numJewels := 0
	for _, unknown := range S {
		if jewelTypes[unknown] {
			numJewels++
		}
	}
	return numJewels
}