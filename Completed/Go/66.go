func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] > 9 {
			carry = 1
			digits[i] = digits[i] % 10
		} else {
			carry = 0
		}
	}

	if carry == 1 {
		return append([]int{carry}, digits...)
	}
	return digits
}