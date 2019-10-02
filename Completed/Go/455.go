import "sort"

func findContentChildren(g []int, s []int) int {
	result := 0
	anchor := 0

	sort.Ints(g)
	sort.Ints(s)

	for _, cookie := range s {
		if anchor > len(g)-1 {
			break
		}

		if cookie >= g[anchor] {
			result++
			anchor++
		}
	}

	return result
}