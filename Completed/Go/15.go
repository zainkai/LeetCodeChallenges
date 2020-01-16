import "sort"

// HARD
/*
If unique results are required, O(N^2) using two pointers is the main optimal solution
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { // avoid duplicate values
			continue
		}

		for j, k := i+1, len(nums)-1; j < k; {
			if j > i+1 && nums[j] == nums[j-1] { // optimization avoiding duplicate nums[j] values to be iterated on
				j++
				continue
			}

			if sum := nums[i] + nums[j] + nums[k]; sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})

				j++ // increment once for next for loop, to avoid duplicates
				for ; j < k && nums[j] == nums[j-1]; j++ {
				}
			} else if sum < 0 { // sum is too low add nums
				j++
			} else { // sum is too high subtract nums
				k--
			}
		}
	}

	return result
}