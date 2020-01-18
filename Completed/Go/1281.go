func subtractProductAndSum(n int) int {
	sum, product := 0, 1

	for n > 0 {
		mod := n % 10
		sum += mod
		product *= mod

		n /= 10
	}

	return product - sum
}