import "math"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	result := 0
	minStock := math.MaxInt64
	maxStock := math.MinInt64

	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > maxStock {
			maxStock = prices[i]
			minStock = math.MaxInt64
		}
		if prices[i] < minStock {
			minStock = prices[i]
		}

		result = max(result, maxStock-minStock)
	}

	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}