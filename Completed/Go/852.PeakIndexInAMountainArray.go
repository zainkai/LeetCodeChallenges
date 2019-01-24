func peakIndexInMountainArray(A []int) int {
	var idx, result int = -1, -1
	for i, val := range A {
			if val > result {
					idx = i
					result = val
			}        
	}
	return idx
}