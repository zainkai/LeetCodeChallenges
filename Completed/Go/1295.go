func findNumbers(nums []int) int {
	countEvens := 0
	for _, num := range nums {
		if numDigits(num)%2 == 0 {
			countEvens++
		}
	}

	return countEvens
}

func numDigits(num int) int {
	i := 0
	for num > 0 {
		num /= 10
		i++
	}

	return i
}