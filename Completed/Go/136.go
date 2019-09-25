func singleNumber(nums []int) int {
	m := map[int]int{}

	for _, num := range nums {
		m[num] += 1
	}

	for key, val := range m {
		if val == 1 {
			return key
		}
	}

	return -1
}

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res ^ v
	}

	return res
}