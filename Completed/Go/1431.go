func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandy := getMax(candies)
    
    results := make([]bool, len(candies))
    for i, candy := range candies {
        if candy+extraCandies >= maxCandy {
            results[i] = true
        }
    }
    
    return results
}

func getMax(arr []int) int {
    m := arr[0]
    for _, v := range arr {
        if v > m {
            m = v
        }
    }
    return m
}
