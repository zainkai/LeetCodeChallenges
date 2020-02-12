import (
	"math"
	"sort"
)

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(a, b int) bool {
		return distFromOrigin(points[a]) < distFromOrigin(points[b])
	})

	vals := [][]int{}
	for i := 0; i < K; i++ {
		vals = append(vals, points[i])
	}

	return vals
}

func distFromOrigin(point []int) float64 {
	x, y := 0-point[0], 0-point[1]
	x *= x
	y *= y

	return math.Sqrt(float64(x + y))
}