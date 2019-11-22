func checkSubarraySum(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}

	arr := preProcessSums(nums)

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[j] - arr[i] + nums[i]
			if isMultipleOfK(sum, k) {
				return true
			}
		}
	}
	return false
}

func isMultipleOfK(x, k int) bool {
	if k == x {
		return true
	}

	return k != 0 && x%k == 0
}

func preProcessSums(nums []int) []int {
	arr := make([]int, len(nums))
	for i := 1; i < len(arr); i++ {
		arr[i] = nums[i] + arr[i-1]
	}

	return arr
}