func isHappy(n int) bool {
	slow := n
	fast := computeHappy(n)
	for slow != fast {
		slow = computeHappy(slow)
		fast = computeHappy(computeHappy(fast))
	}

	return fast == 1
}

func computeHappy(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit

		n /= 10
	}

	return sum
}