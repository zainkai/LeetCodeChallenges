func dailyTemperatures(T []int) []int {
    mapTemps := make(map[int]int)
    res := make([]int, len(T))
    
    for i := len(T)-1 ; i >= 0; i-- {
        t := T[i]
        for temp, pos := range mapTemps {
            if temp > t && res[i] == 0 {
                res[i] = pos-i
            } else if temp > t && res[i] > pos-i {
                res[i] = pos-i
            }
        }
        
        if mapTemps[t] == 0 {
           mapTemps[t] = i 
        } else {
            mapTemps[t] = min(i, mapTemps[t])
        }
    }
    
    return res
    
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}