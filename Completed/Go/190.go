func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	var mask uint32 = 1

	for i := 0; i < 31; i++ {
		if num&mask > 0 {
			result += 1
		}

		result <<= 1
		mask <<= 1
	}

	if num&mask > 0 {
		result += 1
	}

	return result
}