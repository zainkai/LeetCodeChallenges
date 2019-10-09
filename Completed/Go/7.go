import "math"

func reverse(x int) int {
	t := []int{}
	isNeg, y := isNegABS(x)

	for y > 0 {
		t = append(t, y%10)
		y /= 10
	}

	result := 0
	for i, val := range t {
		j := math.Pow(10, float64(len(t)-1-i))
		result += int(j) * val

		if result > math.MaxInt32 {
			return 0
		}
	}

	if isNeg {
		return -result
	}
	return result
}

func isNegABS(x int) (bool, int) {
	if x < 0 {
		return true, -x
	}
	return false, x
}

