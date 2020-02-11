import "math"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return helper([][]int{p1, p2, p3, p4}, 1)
}

func helper(p [][]int, i int) bool {
	if i == len(p) {
		return checkSquare(p[0], p[1], p[2], p[3])
	}
	for j := range p {
		p[i], p[j] = p[j], p[i]
		if helper(p, i+1) {
			return true
		}
		p[i], p[j] = p[j], p[i] // backtrack
	}
	return false
}

func checkSquare(p1, p2, p3, p4 []int) bool {
	a := dist(p1, p2)
	if a == 0 {
		return false
	}

	b := dist(p2, p3)
	c := dist(p3, p4)
	d := dist(p4, p1)

	e := dist(p1, p3)
	f := dist(p2, p4)

	return a == b && b == c && c == d && d == a && e == f
}

func dist(a, b []int) float64 {
	x := (a[0] - b[0])
	x *= x

	y := (a[1] - b[1])
	y *= y

	return math.Sqrt(float64(x + y))
}