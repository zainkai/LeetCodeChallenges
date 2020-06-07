func generate(numRows int) [][]int {
    if numRows == 0 {
        return [][]int{}
    } else if numRows == 1 {
        return [][]int{{1}}
    } else if numRows == 2 {
        return [][]int{{1},{1,1}}
    }
    
    res := [][]int{
        {1},
        {1,1},
    }
    numRows -= 2
    for numRows > 0 {
        topRow := res[len(res)-1]
        
        tmp := []int{1,}
        for i := 0; i < len(topRow)-1;i++ {
            curr := topRow[i]+topRow[i+1]
            tmp = append(tmp, curr)
        }
        tmp = append(tmp, 1)
        
        res = append(res, tmp)
        numRows--
    }
    return res
}
