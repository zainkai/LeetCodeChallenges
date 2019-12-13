import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	newInternals := make([][]int, 0)
	a, b := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		c, d := intervals[i][0], intervals[i][1]
		if c <= b && b < d {
			b = d
		} else if c > b {
			newInternals = append(newInternals, []int{a, b})
			a = c
			b = d
		}
	}
	newInternals = append(newInternals, []int{a, b})
	return newInternals

}