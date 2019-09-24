import "sort"

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}

	// count
	for _, v := range nums {
		m[v] += 1
	}

	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys[:k]
}