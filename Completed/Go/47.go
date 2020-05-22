import (
    "sort"
)

func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    
    results := [][]int{}
    helper(0, nums, &results)
    
    
    return results
}

func helper(idx int, nums []int, results *[][]int) {
    if idx == len(nums) {
        *results = append(*results, deepCopy(nums))
        return
    }
    
    seen := map[int]bool{}
    for i := idx; i < len(nums); i++ {
        num := nums[i]
        if seen[num] {
            continue
        }
        
        nums[idx], nums[i] = nums[i], nums[idx]
        helper(idx+1, nums, results)
        nums[idx], nums[i] = nums[i], nums[idx]
        
        seen[num] = true
    }
}

func deepCopy(nums []int) []int {
    tmp := make([]int, len(nums))
    copy(tmp, nums)
    return tmp
}
