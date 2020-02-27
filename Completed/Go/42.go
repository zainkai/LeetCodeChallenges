func trap(height []int) int {
	leftM := getLeftMax(height)   // O(N)
	rightM := getRightMax(height) //O(N)

	result := 0
	for i := 1; i < len(height)-1; i++ {
		result += min(leftM[i], rightM[i]) - height[i]
	}

	return result
}

func getLeftMax(height []int) []int {
	maxes := make([]int, len(height))
	currMax := 0
	for i, val := range height {
		currMax = max(currMax, val)
		maxes[i] = currMax
	}

	return maxes
}

func getRightMax(height []int) []int {
	maxes := make([]int, len(height))
	currMax := 0
	for i := len(height) - 1; i >= 0; i-- {
		currMax = max(currMax, height[i])
		maxes[i] = currMax
	}

	return maxes
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}