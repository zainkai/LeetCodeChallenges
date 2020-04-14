// TLE O(N^2)
func findMaxLength(nums []int) int {
    maxLen := 0
    for i := 0; i < len(nums); i++ {
        count := 0
        for j := i; j < len(nums); j++ {
            if nums[j]== 1 {
                count += 1
            } else {
                count -= 1
            }
            
            if count == 0 {
                maxLen = max(maxLen, j- i+1)
            }
        }
    }
    
    return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// O(N)
func findMaxLength(nums []int) int {
	idxMap := map[int]int{
		0: -1,
	}
	maxLen, count := 0, 0
	for i, val := range nums {
        if val == 1 {
            count += 1
        } else {
            count -= 1
        }
        
		if j, ok := idxMap[count]; ok {
			maxLen = max(maxLen, i-j)
		} else {
			idxMap[count] = i
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}