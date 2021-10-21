import (
    "sort"
)

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    res := [][]int{}
    sTmp, eTmp := intervals[0][0], intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        if isOverlap([]int{sTmp, eTmp}, intervals[i]) {
            eTmp = max(eTmp, intervals[i][1])
        } else {
            s, e := intervals[i][0], intervals[i][1]
            
            res = append(res, []int{sTmp, eTmp})
            sTmp, eTmp = s, e
        }
    }
    res = append(res, []int{sTmp, eTmp})
    
    return res
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func isOverlap(inv1, inv2 []int) bool {
    s1, e1 := inv1[0], inv1[1]
    s2, _ := inv2[0], inv2[1]
    
    if s2 >= s1 && s2 <= e1 {
        return true
    }
    return false
}
