func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count += 1
		}

		num = num >> 1
	}

	return count
}