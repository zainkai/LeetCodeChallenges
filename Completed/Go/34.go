func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	hi := len(nums) - 1
	lo := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return expandAroundCenter(nums, mid)
		} else if target < nums[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return []int{-1, -1}
}

func expandAroundCenter(nums []int, index int) []int {
	start := index
	end := index

	ele := nums[index]
	for i := index; i >= 0; i-- {
		if nums[i] == ele {
			start = i
		} else {
			break
		}
	}
	for j := index; j < len(nums); j++ {
		if nums[j] == ele {
			end = j
		} else {
			break
		}
	}

	return []int{start, end}
}