func maximumWealth(accounts [][]int) int {
    maxWealth := 0
    for _, banks := range accounts {
        maxWealth = max(maxWealth, sumArr(banks))
    }
    return maxWealth
}

func sumArr (arr []int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}

func max (a,b int) int {
    if a > b {
        return a
    }
    return b
}
