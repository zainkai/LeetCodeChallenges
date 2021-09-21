func maxProfit(prices []int) int {
    if len(prices) <= 1 {
        return 0
    }
    
    maxSell := 0
    savedSell := make([]int, len(prices))
    for i := len(savedSell)-2; i >= 0; i-- {
        maxSell = max(maxSell, prices[i+1])
        savedSell[i] = maxSell
    }
    
    maxProfit := 0
    for i, price := range prices {
        tprofit := savedSell[i] - price
        maxProfit = max(maxProfit, tprofit)
    }
    
    return maxProfit
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
