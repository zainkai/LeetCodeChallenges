import (
	"sort"
)

func highFive(items [][]int) [][]int {
	m := map[int][]int{}
	for _, n := range items {
		id, score := n[0], n[1]
		m[id] = append(m[id], score)
	}

	res := make([][]int, len(m))
	for id, scores := range m {
		sort.Slice(scores, func(i, j int) bool {
			return scores[i] > scores[j]
		})
		res[id-1] = []int{id, getAvg(scores[:5])}
	}

	return res
}

func getAvg(scores []int) int {
	sum := 0
	for _, v := range scores {
		sum += v
	}

	return sum / len(scores)
}