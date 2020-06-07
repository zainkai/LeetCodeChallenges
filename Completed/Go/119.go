func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    
    res := []int{1,1}
    rowIndex -= 1
    for rowIndex > 0 {
        tmp := []int{1,}
        for i := 0; i < len(res)-1; i++ {
            curr := res[i]+res[i+1]
            tmp = append(tmp, curr)
        }
        res = append(tmp, 1)
        
        rowIndex--
    }
    return res
}
