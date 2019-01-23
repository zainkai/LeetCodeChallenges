func sortArrayByParity(A []int) []int {
	evens := make([]int, 0)
	odds := make([]int, 0)
	
	for _,a := range A {
			if a % 2 == 0 {
					evens = append(evens, a)
			} else {
					odds = append(odds, a)
			}
	}
	return append(evens, odds...)
}