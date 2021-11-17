func getModifiedArray(length int, updates [][]int) []int {
    res := make([]int, length)
    startMap := map[int]int{}
    endMap := map[int]int{}
    
    for _, update := range updates {
        startIdx, endIdx, inc := update[0], update[1], update[2]
        
        startMap[startIdx] += inc
        endMap[endIdx] -= inc
    }
        
    anchor := 0
    for i := range res {
        anchor += startMap[i]
        res[i] = anchor
        anchor += endMap[i]
    }
 
    return res
}
