func smallestCommonElement(mat [][]int) int {
    m := map[int]int{}
    
    for _, arr := range mat {
        for _, v := range arr {
            m[v] += 1
        }
    }
    
    for _, key := range getSortedKeys(m) {
        if m[key] == len(mat) {
            return key
        }
    }
    
    return -1
}

func getSortedKeys(m map[int]int) []int {
    result := []int{}
    
    for k := range m {
        result = append(result, k)
    }
    
    sort.Ints(result)
    return result
}