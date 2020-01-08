package notes

func search(nums []int, t int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == t {
			return mid
		} else if t < nums[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
