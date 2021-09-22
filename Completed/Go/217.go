// time O(n) space O(n)

func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}

	for _, key := range nums {
		if _, ok := m[key]; ok {
			return true
		}
		m[key] = struct{}{}
	}

	return false
}

// time O(n log n) space O(1)
import "sort"

func containsDuplicate(nums []int) bool {
    seenDup := false
    sort.Slice(nums, func(i, j int) bool {
        a, b := nums[i], nums[j]
        if a == b {
            seenDup = true
        }
        
        if a < b {
            return true
        } else {
            return false
        }
    })
    
    return seenDup
}
