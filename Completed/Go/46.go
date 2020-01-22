// First Solution
// func permute(nums []int) [][]int {
//   res := make([][]int, 0)
//   helper(nums, []int{}, &res)

//   return res
// }

// func helper(nums []int, output []int, res *[][]int) {
//   if len(nums) == 0 {
//     *res = append(*res, output)
//   }

//   for i, val := range nums {
//     tempNums := make([]int, 0)
//     tempNums = append(tempNums, nums[:i]...)
//     tempNums = append(tempNums, nums[i+1:]...)

//     tempOutput := make([]int, len(output))
//     copy(tempOutput, output)
//     tempOutput = append(tempOutput, val)

//     helper(tempNums, tempOutput, res)
//   }
// }

// Better Solution 2nd iteration
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	helper(0, nums, &res)

	return res
}

func helper(i int, nums []int, res *[][]int) {
	if i == len(nums)-1 {
		temp := make([]int, len(nums))
		copy(temp, nums)

		*res = append(*res, temp)
		return
	}

	for j := i; j < len(nums); j++ {
		nums[i], nums[j] = nums[j], nums[i]
		helper(i+1, nums, res)
		nums[j], nums[i] = nums[i], nums[j]
	}
}
