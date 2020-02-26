func findDuplicates(nums []int) []int {
	dups := []int{}

	for _, num := range nums {
		if i := getIndex(num); nums[i] < 0 {
			dups = append(dups, i+1)
		} else {
			nums[i] = -nums[i]
		}
	}

	return dups
}

func getIndex(num int) int {
	if num < 0 {
		num = (num * -1)
	}
	return num - 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}