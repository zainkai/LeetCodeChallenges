func hasAlternatingBits(n int) bool {
	for n > 0 {
		isFirstBitSet := n&1 > 0
		isSecondBitSet := n&2 > 0

		if isFirstBitSet != isSecondBitSet {
			n >>= 1
		} else {
			return false
		}
	}

	return true
}