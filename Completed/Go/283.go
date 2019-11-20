// O(N^2)
// func moveZeroes(nums []int) {
// 	end := 0

// 	for i := 0; i < len(nums)-1-end; i++ {
// 		for j := i + 1; j < len(nums)-end; j++ {
// 			if nums[i] == 0 {
// 				swap(nums, i, j)

// 			} else if nums[i] != 0 && nums[j] != 0 {
// 				break
// 			}

// 			if j == len(nums)-1-end {
// 				end++
// 			}
// 		}
// 	}
// }

func moveZeroes(nums []int) {
	lastKnownZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			swap(nums, i, lastKnownZero)
			lastKnownZero++
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}