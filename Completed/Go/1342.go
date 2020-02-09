func numberOfSteps(num int) int {
	i := 0
	for num > 0 {
		i++
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
	}

	return i
}