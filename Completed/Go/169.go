
//O(N) space and time
func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num] += 1

		if m[num] > len(nums)/2 {
			return num
		}
	}

	return -1
}

// O(N) time

import (
	"sort"
)

func majorityElement(nums []int) int {
	mid := int(len(nums)/2)
	sort.Ints(nums)
	
	return nums[mid]
}


